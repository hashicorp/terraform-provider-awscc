// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

// Legacy code: tests the legacy schema generator, superseded by bigdiffer's
// discovery + `-update`/`-generate`/`-docs` pipeline. Kept as the deprecated
// make/go:generate fallback; see
// contributing/docs/generating-the-provider-with-bigdiffer.md#fallback-the-legacy-process.

package main

import (
	"testing"

	"github.com/hashicorp/cli"
)

func TestDownloaderMetaSchema(t *testing.T) {
	d := &Downloader{
		config: Config{
			MetaSchema: MetaSchema{
				Path: "../../../service/cloudformation/meta-schemas/provider.definition.schema.v1.json",
			},
		},
		ui: cli.NewMockUi(),
	}

	err := d.MetaSchema()

	if err != nil {
		t.Fatalf("Downloader.MetaSchema: %s", err)
	}
}
