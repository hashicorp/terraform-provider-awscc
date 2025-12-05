// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
)

type waiterResultError struct {
	errorCode  string
	errorMsg   string
	hookErrors []string
}

func (e *waiterResultError) Error() string {
	if e.errorCode != "" {
		e.errorMsg += fmt.Sprintf(". ErrorCode: %s", e.errorCode)
	}

	if len(e.hookErrors) > 0 {
		e.errorMsg += ". Hook failures: " + strings.Join(e.hookErrors, "; ")
	}

	return e.errorMsg
}

func newWaiterErr(status, msg string) *waiterResultError {
	e := waiterResultError{}
	e.errorMsg = fmt.Sprintf("waiter state transitioned to %s. StatusMessage: %s", status, msg)

	return &e
}

func (e *waiterResultError) withErrorCode(code string) {
	e.errorCode = code
}

func (e *waiterResultError) withHookEvent(hookEvent types.HookProgressEvent) {
	e.hookErrors = append(e.hookErrors, fmt.Sprintf("HookName: %s, HookArn: %s, HookVersion: %s, Time: %s, HookMessage: %s",
		aws.ToString(hookEvent.HookTypeName),
		aws.ToString(hookEvent.HookTypeArn),
		aws.ToString(hookEvent.HookTypeVersionId),
		hookEvent.HookEventTime.Format(time.RFC3339),
		aws.ToString(hookEvent.HookStatusMessage)))
}
