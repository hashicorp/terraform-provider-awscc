//go:build !writeerrors

// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

// writeProviderError writes a factory error to the log file.
// kind should be "resource", "data source", or "list resource".
func writeProviderError(_ string, _ string, _ error) {}
