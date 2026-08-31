// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Parallel()

	cfg, err := newConfig(
		"internal/provider/all_schemas.hcl",
		defaultsBlock{
			SchemaCacheDirectory:    "../service/cloudformation/schemas",
			TerraformTypeNamePrefix: "awscc",
		},
		metaSchemaBlock{Path: "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"},
	)
	if err != nil {
		t.Fatalf("newConfig: %v", err)
	}

	// Every path must be absolute so nothing depends on the working directory.
	for name, p := range map[string]string{
		"overlayPath":      cfg.overlayPath,
		"overlayDir":       cfg.overlayDir,
		"cacheDir":         cfg.cacheDir,
		"metaSchemaPath":   cfg.metaSchemaPath,
		"servicesPath":     cfg.servicesPath,
		"outputRoot":       cfg.outputRoot,
		"registrationPath": cfg.registrationPath,
	} {
		if !filepath.IsAbs(p) {
			t.Errorf("%s is not absolute: %s", name, p)
		}
	}

	// Relative-to-overlay paths resolve as the legacy generators expected.
	wantSuffix := map[string]string{
		"cacheDir":       filepath.Join("internal", "service", "cloudformation", "schemas"),
		"metaSchemaPath": filepath.Join("internal", "service", "cloudformation", "meta-schemas", "provider.definition.schema.v1.json"),
		"servicesPath":   filepath.Join("internal", "identity", "names", "services.hcl"),
		"outputRoot":     "internal",
		"overlayDir":     filepath.Join("internal", "provider"),
	}
	got := map[string]string{
		"cacheDir":       cfg.cacheDir,
		"metaSchemaPath": cfg.metaSchemaPath,
		"servicesPath":   cfg.servicesPath,
		"outputRoot":     cfg.outputRoot,
		"overlayDir":     cfg.overlayDir,
	}
	for name, suffix := range wantSuffix {
		if !strings.HasSuffix(got[name], suffix) {
			t.Errorf("%s = %q, want suffix %q", name, got[name], suffix)
		}
	}

	if cfg.prefix != "awscc" {
		t.Errorf("prefix = %q, want awscc", cfg.prefix)
	}
	if cfg.region != discoverRegion {
		t.Errorf("region = %q, want %q", cfg.region, discoverRegion)
	}
	if cfg.genConcurrency < 1 {
		t.Errorf("genConcurrency = %d, want >= 1", cfg.genConcurrency)
	}
}
