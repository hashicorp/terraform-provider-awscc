package rename

import (
    "encoding/json"
    "io"
    "os"
    "strings"
    "reflect"
    "github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
)

type Generator struct {
	*common.Generator
}

func NewGenerator() *Generator {
	return &Generator{
		Generator: common.NewGenerator(),
	}
}

const filePermMode = 0644 // Read/write for owner, read for group/others

func (g *Generator) RenameCfnSchemaFile (filePath string) error {
    // Open JSON schema file from base directory 
    file, err := os.Open(filePath)
    if err != nil {
        g.Errorf("Error opening file:", err)
        return err
    }
    defer file.Close()

    // Read the JSON data
    data, err := io.ReadAll(file)
    if err != nil {
        g.Errorf("Error reading file:", err)
        return err
    }

    // Unmarshal the JSON data into a map while preserving the order
    var jsonData map[string]interface{}
    err = json.Unmarshal(data, &jsonData)
    if err != nil {
        g.Errorf("Error unmarshaling JSON:", err)
        return err
    }

    // Create a copy of the original JSON data
    originalData := make(map[string]interface{})
    for k, v := range jsonData {
        originalData[k] = v
    }

    // Replace "CloudFormation" with "Terraform" in the description
    err = updateDescription(jsonData)
    if err != nil {
        g.Errorf("Error updating description:", err)
        return err
    }

    // Check if the JSON data has changed
    if reflect.DeepEqual(jsonData, originalData) {
        // No changes detected, skip writing the file
        return nil
    }

    // Marshal the updated JSON data while preserving the order
    updatedDataBytes, err := json.MarshalIndent(jsonData, "", "  ")
    if err != nil {
        g.Errorf("Error marshaling JSON:", err)
        return err
    }

    // Write the updated JSON data back to the file
    err = os.WriteFile(filePath, updatedDataBytes, filePermMode)
    if err != nil {
        g.Errorf("Error writing file:", err)
        return err
    }

    g.Infof("File %s updated successfully\n", filePath)
    return nil
}

func updateDescription(data map[string]interface{}) error {
    for key, value := range data {
        if key == "description" {
            description, ok := value.(string)
            if ok {
                updatedDescription := strings.ReplaceAll(description, "AWS CloudFormation", "Terraform")
                data[key] = updatedDescription
            }
        } else {
            switch v := value.(type) {
            case map[string]interface{}:
                err := updateDescription(v)
                if err != nil {
                    return err
                }
            case []interface{}:
                for i, item := range v {
                    switch item := item.(type) {
                        case map[string]interface{}:
                            err := updateDescription(item)
                            if err != nil {
                                return err
                            }
                            v[i] = item
                    }
                }
            }
        }
    }
    return nil
}