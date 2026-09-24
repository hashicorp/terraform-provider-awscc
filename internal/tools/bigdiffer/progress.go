// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/term"
)

// barThrottle bounds how often the progress bar re-renders.
const barThrottle = 100 * time.Millisecond

// stepf prints a top-level step header to stderr (e.g. "==> Discovering…"), so a
// maintainer can follow what bigdiffer is doing at each stage.
func stepf(format string, a ...any) {
	fmt.Fprintf(structOut, "==> "+format+"\n", a...)
}

// infof prints an indented detail line beneath the current step.
func infof(format string, a ...any) {
	fmt.Fprintf(structOut, "    "+format+"\n", a...)
}

// newBar returns a labelled progress bar on stderr. When stderr is not a terminal
// (CI, redirected logs), it returns a silent bar so logs stay clean. Add and
// Finish are safe to call concurrently.
func newBar(total int, label string) *progressbar.ProgressBar {
	if total <= 0 || !term.IsTerminal(int(os.Stderr.Fd())) {
		return progressbar.DefaultSilent(int64(total), label)
	}
	return progressbar.NewOptions(total,
		progressbar.OptionSetDescription(label),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowCount(),
		progressbar.OptionShowElapsedTimeOnFinish(),
		// Predicted "time remaining" is meaningless for this workload: artifacts
		// vary in cost and 16 workers saturate after a slow start, so a linear
		// ETA drifts up early then lurches down. Show elapsed time (which a
		// determinate bar hides by default) with no remaining estimate.
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionThrottle(barThrottle),
		progressbar.OptionClearOnFinish(),
	)
}

// newToolBar returns an indeterminate console spinner for an external tool
// (terraform fmt, tfplugindocs) whose per-line output is being routed to the
// run log instead of the console. It shows a live processed-line count and
// elapsed time; there is no honest denominator for these tools, so it is a
// spinner rather than a filling bar. Silent (no spinner) when stderr is not a
// terminal, so CI logs stay clean.
func newToolBar(label string) *progressbar.ProgressBar {
	if !term.IsTerminal(int(os.Stderr.Fd())) {
		return progressbar.DefaultSilent(-1, label)
	}
	return progressbar.NewOptions(-1,
		progressbar.OptionSetDescription(label),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionShowIts(),
		progressbar.OptionSetItsString("lines"),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionThrottle(barThrottle),
		progressbar.OptionClearOnFinish(),
	)
}
