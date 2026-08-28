// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// config is bigdiffer's single source of truth for the paths and tunables the
// pipeline needs. It is derived once from the overlay (its defaults/meta blocks)
// plus flags and constants, with every relative-to-overlay path resolved to an
// absolute path here — so nothing downstream depends on the current working
// directory (the brittle assumption the legacy generators and our own early
// gate tests had to work around).
type config struct {
	overlayPath        string // all_schemas.hcl (absolute)
	overlayDir         string // directory containing the overlay (…/internal/provider)
	cacheDir           string // committed CloudFormation schema JSONs
	metaSchemaPath     string // provider meta-schema JSON
	servicesPath       string // identity/names/services.hcl (IsGlobal/HasMutableIdentity)
	outputRoot         string // base dir for generated <svc>/… code (…/internal)
	registrationPath   string // the single blank-import registration file bigdiffer emits
	importExamplesPath string // import_examples_gen.json aggregate
	repoRoot           string // repository root (…/, parent of internal)
	examplesDir        string // examples/ (import-example docs land here)
	docsDir            string // docs/ (tfplugindocs output)

	prefix         string // terraform_type_name_prefix, e.g. "awscc"
	region         string // CloudFormation registry region for -discover
	genConcurrency int    // bounded worker count for in-process generation
}

// newConfig resolves configuration from the overlay path and its decoded
// defaults/meta blocks. All paths are made absolute relative to the overlay's
// directory, matching how the legacy generators resolved them from
// internal/provider (e.g. schema_cache_directory "../service/cloudformation/schemas").
func newConfig(overlayPath string, defaults defaultsBlock, meta metaSchemaBlock) (config, error) {
	absOverlay, err := filepath.Abs(overlayPath)
	if err != nil {
		return config{}, fmt.Errorf("resolving overlay path %q: %w", overlayPath, err)
	}
	overlayDir := filepath.Dir(absOverlay)
	outputRoot := filepath.Dir(overlayDir) // …/internal/provider -> …/internal

	return config{
		overlayPath:        absOverlay,
		overlayDir:         overlayDir,
		cacheDir:           filepath.Clean(filepath.Join(overlayDir, defaults.SchemaCacheDirectory)),
		metaSchemaPath:     filepath.Clean(filepath.Join(overlayDir, meta.Path)),
		servicesPath:       filepath.Join(outputRoot, "identity", "names", "services.hcl"),
		outputRoot:         outputRoot,
		registrationPath:   filepath.Join(overlayDir, "registrations_gen.go"),
		importExamplesPath: filepath.Join(overlayDir, "import_examples_gen.json"),
		repoRoot:           filepath.Dir(outputRoot),
		examplesDir:        filepath.Join(filepath.Dir(outputRoot), "examples"),
		docsDir:            filepath.Join(filepath.Dir(outputRoot), "docs"),
		prefix:             defaults.TerraformTypeNamePrefix,
		region:             discoverRegion,
		genConcurrency:     runtime.NumCPU(),
	}, nil
}

// loadOverlay decodes the overlay file and builds the config from its
// defaults/meta blocks, returning the config and the resource rows.
func loadOverlay(overlayPath string) (config, []resourceRow, error) {
	var f allSchemasFile
	if err := hclsimple.DecodeFile(overlayPath, nil, &f); err != nil {
		return config{}, nil, fmt.Errorf("decoding overlay %s: %w", overlayPath, err)
	}
	cfg, err := newConfig(overlayPath, f.Defaults, f.Meta)
	if err != nil {
		return config{}, nil, err
	}
	return cfg, f.Resources, nil
}
