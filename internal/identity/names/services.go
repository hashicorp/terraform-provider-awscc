// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package names

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	Services []Service `hcl:"service,block"`
}

type Service struct {
	IsGlobal    bool       `hcl:"is_global,optional"`
	ServiceName string     `hcl:",label"`
	Resources   []Resource `hcl:"resource,block"`
}

type Resource struct {
	TFResourceName     string `hcl:",label"`
	HasMutableIdentity bool   `hcl:"has_mutable_identity,optional"`
}

func ParseServicesFile(filename string) (*Config, error) {
	var config Config
	err := hclsimple.DecodeFile(filename, nil, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GetServiceName(s string) string {
	parts := strings.Split(s, "::")
	if len(parts) > 1 {
		return strings.ToLower(parts[1])
	}

	return ""
}
