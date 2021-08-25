package tflog

import (
	"io"
	"os"
)

// loggerKey defines context keys for locating loggers in context.Context
// it's a private type to make sure no other packages can override the key
type loggerKey string

const (
	providerSpaceRootLoggerKey loggerKey = "provider"
)

var (
	// When go-plugin.Serve is called, it overwrites our os.Stderr with a
	// gRPC stream which Terraform ignores. This tends to be before our
	// loggers get set up, as go-plugin has no way to pass in a base
	// context, and our loggers are passed around via contexts. This leaves
	// our loggers writing to an output that is never read by anything,
	// meaning the logs get blackholed. This isn't ideal, for log output,
	// so this is our workaround: we copy stderr on init, before Serve can
	// be called, and offer an option to write to that instead of the
	// os.Stderr available at runtime.
	//
	// Ideally, this is a short-term fix until Terraform starts reading
	// from go-plugin's gRPC-streamed stderr channel, but for the moment it
	// works.
	stderr io.Writer
)

func init() {
	stderr = os.Stderr
}
