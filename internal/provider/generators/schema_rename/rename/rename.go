package rename

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "reflect"
)

func RenameCfnSchemaFile (filePath string) error {
    // Open JSON schema file from base directory 
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return err
    }
    defer file.Close()

    // Read the JSON data
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return err
    }

    // Unmarshal the JSON data into a map while preserving the order
    var jsonData map[string]interface{}
    err = json.Unmarshal(data, &jsonData)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
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
        fmt.Println("Error updating description:", err)
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
        fmt.Println("Error marshaling JSON:", err)
        return err
    }

    // Write the updated JSON data back to the file
    err = ioutil.WriteFile(filePath, updatedDataBytes, 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return err
    }

    fmt.Printf("File %s updated successfully\n", filePath)
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
                    switch item.(type) {
                    case map[string]interface{}:
                        err := updateDescription(item.(map[string]interface{}))
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