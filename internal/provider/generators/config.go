package generators

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// Resource schema configuration file format.

type Config struct {
	MetaSchema      MetaSchema       `hcl:"meta_schema,block"`
	ResourceSchemas []ResourceSchema `hcl:"resource_schema,block"`
}

type MetaSchema struct {
	Local   string `hcl:"local"`
	Refresh bool   `hcl:"refresh,optional"`
	Source  Source `hcl:"source,block"`
}

type ResourceSchema struct {
	Local        string `hcl:"local"`
	Refresh      bool   `hcl:"refresh,optional"`
	ResourceName string `hcl:"resource_name,label"`
	Source       Source `hcl:"source,block"`
}

type Source struct {
	Url string `hcl:"url"`
}

// NewConfigPath returns a Config or any errors from the provided HCL at the file path.
func NewConfigPath(path string) (*Config, error) {
	var config Config

	if err := hclsimple.DecodeFile(path, nil, &config); err != nil {
		return nil, fmt.Errorf("error decoding configuration: %w", err)
	}

	return &config, nil
}
