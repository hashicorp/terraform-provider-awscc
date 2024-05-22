// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
