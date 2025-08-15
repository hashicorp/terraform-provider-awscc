package main

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func main() {

	out, err := ParseServicesFile()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(out.Services[0].Resources)
}

type Config struct {
	Services []Service `hcl:"service,block"`
}

type Service struct {
	IsGlobal    bool       `hcl:"is_global,optional"`
	ServiceName string     `hcl:",label"`
	Resources   []Resource `hcl:"resource,block"`
}

type Resource struct {
	TFResourceName     string `hcl:"tf_resource_name"`
	HasMutableIdentity bool   `hcl:"has_mutable_identity,optional"`
}

func ParseServicesFile() (*Config, error) {
	var config Config
	err := hclsimple.DecodeFile("services.hcl", nil, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
