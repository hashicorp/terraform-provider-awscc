//go:build !writeerrors

// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

// writeResourceError is a no-op in normal builds.
func writeResourceError(_ string, _ error) {}
