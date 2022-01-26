package generic

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	LoggingKeyCFNType = "cfn_type"
)

func traceEntry(ctx context.Context, n string) {
	tflog.Trace(ctx, fmt.Sprintf("%s entry", n))
}

func traceExit(ctx context.Context, n string) {
	tflog.Trace(ctx, fmt.Sprintf("%s exit", n))
}
