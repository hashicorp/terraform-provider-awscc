// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"errors"
	"testing"
)

func TestDecide(t *testing.T) {
	t.Parallel()

	const today = "2026-08-27"

	okRes := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
		{kind: artifactResource, outcome: gateOK},
		{kind: artifactSingularDataSource, outcome: gateOK},
	}}
	brokenRes := gateResult{cfType: "AWS::Svc::Thing", artifacts: []artifactResult{
		{kind: artifactResource, outcome: gateFailedGeneration, err: errors.New("recursive definition\nmore detail")},
		{kind: artifactSingularDataSource, outcome: gateOK},
	}}

	t.Run("new + ok adds a plain block", func(t *testing.T) {
		t.Parallel()
		d := decide(classNew, okRes, today)
		if !d.addBlock || len(d.setAttrs) != 0 {
			t.Errorf("got %+v, want addBlock with no attrs", d)
		}
	})

	t.Run("new + failed adds suppressed with reason", func(t *testing.T) {
		t.Parallel()
		d := decide(classNew, brokenRes, today)
		if !d.addBlock {
			t.Errorf("should still add the block (backlog)")
		}
		if d.setAttrs[attrSuppressResource] != "true" {
			t.Errorf("resource should be suppressed, got %+v", d.setAttrs)
		}
		if _, ok := d.setAttrs[attrSuppressSingular]; ok {
			t.Errorf("singular gated OK, should not be suppressed")
		}
		if d.reason == "" || d.reason != "resource failed-generation: recursive definition" {
			t.Errorf("reason = %q, want first-line failure detail", d.reason)
		}
	})

	t.Run("present + ok keeps the block untouched", func(t *testing.T) {
		t.Parallel()
		d := decide(classPresent, okRes, today)
		if d.addBlock || len(d.setAttrs) != 0 {
			t.Errorf("got %+v, want no add and no attrs", d)
		}
	})

	t.Run("present + failed freezes at last-good", func(t *testing.T) {
		t.Parallel()
		d := decide(classPresent, brokenRes, today)
		if d.addBlock {
			t.Errorf("present must not add a block")
		}
		if d.setAttrs[attrFrozenSince] != today {
			t.Errorf("should freeze at today, got %+v", d.setAttrs)
		}
		if d.reason == "" {
			t.Errorf("freeze should record a reason")
		}
	})

	t.Run("non-provisionable annotates", func(t *testing.T) {
		t.Parallel()
		d := decide(classNonProvisionable, gateResult{}, today)
		if d.setAttrs[attrNonProvisionable] != "true" || d.addBlock {
			t.Errorf("got %+v, want non_provisionable=true, no add", d)
		}
	})

	t.Run("withdrawn freezes", func(t *testing.T) {
		t.Parallel()
		d := decide(classWithdrawn, gateResult{}, today)
		if d.setAttrs[attrFrozenSince] != today || d.addBlock {
			t.Errorf("got %+v, want frozen_since=today, no add", d)
		}
	})
}
