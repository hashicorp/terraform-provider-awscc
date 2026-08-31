// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io"

	"github.com/hashicorp/cli"
	"github.com/hashicorp/terraform-provider-awscc/internal/tools/bigdiffer/codegen"
	"golang.org/x/sync/errgroup"
)

// genResult is one artifact's generated output (or the error that produced none).
// It carries the plan and artifact so callers can locate the target file, fall
// back to the legacy generator, etc.
type genResult struct {
	p    plan
	a    genArtifact
	code []byte
	test []byte
	err  error
}

// generateArtifact renders one artifact's code + test via the owned engine.
func generateArtifact(ui cli.Ui, cfg config, p plan, a genArtifact) (code, test []byte, err error) {
	switch a.kind {
	case artifactResource:
		return codegen.GenerateResource(ui, p.schemaFile, a.tfType, a.packageName, cfg.servicesPath, a.listResource)
	case artifactSingularDataSource:
		return codegen.GenerateSingularDataSource(ui, p.schemaFile, a.tfType, a.packageName, cfg.servicesPath)
	case artifactPluralDataSource:
		return codegen.GeneratePluralDataSource(p.cfType, a.tfType, a.packageName)
	}
	return nil, nil, fmt.Errorf("unknown artifact kind %q", a.kind)
}

// generateCorpus generates every artifact for the given rows concurrently,
// bounded by concurrency workers. onArtifact, when non-nil, is invoked once per
// generated artifact (safe for concurrent use) so callers can drive a progress
// bar. Generation is CPU-bound and independent per artifact, so this
// parallelizes cleanly: each work item writes its own result slot (no shared
// mutable state), the emitter has no package-level globals, and bigdiffer's
// naming serializes the one non-thread-safe dependency (inflection). One
// artifact's failure is captured in-band and never aborts the others.
func generateCorpus(cfg config, rows []resourceRow, concurrency int, onArtifact func()) []genResult {
	if concurrency < 1 {
		concurrency = 1
	}

	type work struct {
		p plan
		a genArtifact
	}
	var items []work
	var planErrs []genResult
	for _, row := range rows {
		p, err := generationPlan(row, cfg.prefix, cfg.cacheDir)
		if err != nil {
			planErrs = append(planErrs, genResult{
				p:   plan{cfType: row.CloudFormationTypeName},
				err: fmt.Errorf("plan: %w", err),
			})
			continue
		}
		for _, a := range p.artifacts {
			items = append(items, work{p: p, a: a})
		}
	}

	results := make([]genResult, len(items))
	ui := &cli.BasicUi{Writer: io.Discard, ErrorWriter: io.Discard}
	var g errgroup.Group
	g.SetLimit(concurrency)
	for i, w := range items {
		g.Go(func() error {
			code, test, err := generateArtifact(ui, cfg, w.p, w.a)
			results[i] = genResult{p: w.p, a: w.a, code: code, test: test, err: err}
			if onArtifact != nil {
				onArtifact()
			}
			return nil
		})
	}
	_ = g.Wait() // per-artifact errors are in-band in each genResult.

	return append(results, planErrs...)
}
