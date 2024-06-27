package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/schema_rename/rename"
)

func main() {
	baseDir := "./internal/service/cloudformation/schemas/"
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			generator := rename.NewGenerator()
			// Skip files that start with "AWS_CloudFormation_"
			if !strings.HasPrefix(info.Name(), "AWS_CloudFormation_") {
				err = generator.RenameCfnSchemaFile(path)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
}
