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

// init opens the resource error log when the provider binary starts.
// The log path is read from $RESOURCE_ERROR_LOG; defaults to .startup_errors.log
// in the current working directory. Uses O_APPEND because terraform invokes the
// provider binary multiple times per plan — O_TRUNC would wipe earlier entries.
// Truncate the file once before running terraform (e.g. in the make target).
// Enabled via the `writeerrors` build tag.
func init() {
	path := os.Getenv("RESOURCE_ERROR_LOG")
	if path == "" {
		path = ".startup_errors.log"
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return
	}
	resourceErrorLog = f
}

// writeProviderError writes a factory error to the log file.
// kind should be "resource", "data source", or "list resource".
func writeProviderError(kind string, name string, err error) {
	if resourceErrorLog == nil {
		return
	}
	fmt.Fprintf(resourceErrorLog, "kind=%s name=%s error=%s\n", kind, name, err)
}
