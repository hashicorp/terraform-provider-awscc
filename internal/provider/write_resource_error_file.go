//go:build writeerrors

// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
)

// resourceErrorLog holds the log file opened at program startup.
var resourceErrorLog *os.File

// init opens and truncates resource_errors.log when the provider binary starts.
// Enabled via the `writeerrors` build tag.
func init() {
	f, err := os.OpenFile("resource_errors.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return
	}
	resourceErrorLog = f
}

// writeResourceError writes a resource factory error to the log file.
func writeResourceError(name string, err error) {
	if resourceErrorLog == nil {
		return
	}
	fmt.Fprintf(resourceErrorLog, "resource=%s error=%s\n", name, err)
}
