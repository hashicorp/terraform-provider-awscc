// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/term"
)

// stepf prints a top-level step header to stderr (e.g. "==> Discovering…"), so a
// maintainer can follow what bigdiffer is doing at each stage.
func stepf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "==> "+format+"\n", a...)
}

// infof prints an indented detail line beneath the current step.
func infof(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "    "+format+"\n", a...)
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
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionThrottle(100*time.Millisecond),
		progressbar.OptionClearOnFinish(),
	)
}
