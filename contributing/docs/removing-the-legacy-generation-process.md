<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Removing the legacy generation process

A checklist for punchlist item D ("Adopt — phase 2: delete the redundant legacy
machinery"). Follow this once bigdiffer has been the weekly driver for several
clean release cycles and nobody has needed the
[legacy fallback](generating-the-provider-with-bigdiffer.md#fallback-the-legacy-process).

This is a **transient tracking doc**, like `generation-punchlist.md` — delete it
once the checklist below is done.

## Before you start

Confirm both of these, in order. Stop and do not proceed if either fails.

1. **Several real weekly cycles have run on bigdiffer with no fallback to the
   legacy process.** "Several" is a judgment call, not a fixed number — the
   point is confidence, not a calendar date. If anyone reached for
   `generating-the-provider.md` in that window, the fallback is still needed;
   wait longer.
2. `go test ./internal/tools/bigdiffer/... -run TestFullCorpusParity` passes
   (0 drift against the full corpus). This is the single test that proves
   bigdiffer's owned generator engine is still byte-identical to the legacy
   one — do not remove the legacy engine while this test would have caught a
   real divergence.

## The checklist

Do these in order. Each step is independently verifiable — build and test after
every step, not just at the end.

1. **Delete the legacy generator engine and directive files.**
   - Delete the legacy resource/data-source generator code (the pre-bigdiffer
     engine `internal/tools/bigdiffer/codegen` was copied from — see
     `bigdiffer-design.md` §10 for the full list of what bigdiffer subsumed).
   - Delete `internal/provider/resources.go`, `internal/provider/singular_data_sources.go`,
     `internal/provider/plural_data_sources.go` (the three legacy blank-import
     directive files). `internal/provider/registrations_gen.go` is bigdiffer's
     replacement and already carries every registration.
   - Delete the `//go:generate` directives that drove the legacy engine (they
     live in the same three files being deleted, plus scattered
     `generators/resource/main.go` / `generators/*-data-source/main.go`
     invocations — grep for `go:generate.*generators/` to find every site).
   - Build (`make build`) and run the full test suite. Nothing should reference
     the deleted packages; if something does, stop and re-scope before
     continuing — do not comment out the reference and move on.

2. **Delete the legacy `make` targets.**
   - Remove the targets `generating-the-provider.md` documents as the legacy
     process: `cleanschemas`, `suppressions`, `schemas`, `commitrefresh`,
     `biglister`, `bigdiffer` (the legacy dated-snapshot-diff target — not
     `internal/tools/bigdiffer`), `commitschemas`, `resources`,
     `commitresources`, `singular-data-sources`, `plural-data-sources`,
     `commitdatas`, `docs-all` (if `docs-all` is legacy-specific; keep it if
     anything non-legacy still depends on it — check first), and any other
     target whose only caller was the legacy process.
   - Grep the `GNUmakefile` for any of the deleted targets as a dependency of a
     target you're keeping, and fix those dependency chains.

3. **Delete `generating-the-provider.md`** (the legacy runbook itself — it
   describes a process that no longer exists).

4. **Update `generating-the-provider-with-bigdiffer.md`:**
   - Delete the entire "Fallback: the legacy process" section — there is no
     fallback anymore.
   - Delete the "What bigdiffer replaces" table's framing as a *replacement*
     (there is nothing left to compare against); fold anything still useful
     into plain prose, or delete the section outright if it no longer adds
     value once there is no "before."
   - Remove the `> [!NOTE]` callout about `registrations_gen.go` superseding
     the three legacy directive files — they are gone, so "supersedes" and
     "during the transition both may be present" are no longer true. Replace
     with a one-line statement that `registrations_gen.go` is the sole
     registration mechanism.
   - Rename the doc (drop "with bigdiffer" — there is no other way to generate
     the provider anymore) if that reads better standalone; not required, use
     judgment.

5. **Tighten `internal/tools/bigdiffer/main.go`'s `checkRegistrationUpToDate`.**
   This is the one **behavior change**, not just a deletion — easy to forget
   since nothing will fail loudly if you skip it.
   - Today, `checkRegistrationUpToDate` treats a **missing**
     `registrations_gen.go` as `""` (no problem) — see the function's own
     doc comment: *"Absence is not a failure: while the legacy directive files
     ... still register every type, this file is additive ... Once the legacy
     files are removed, presence should be required here."* That condition is
     now true.
   - Change the `errors.Is(err, os.ErrNotExist)` branch to return a problem
     string instead of `""`, so `-check` fails if `registrations_gen.go` is
     ever missing or deleted by hand once it is the *only* registration
     mechanism.
   - Update the function's doc comment to remove the now-stale "Absence is not
     a failure" paragraph.
   - Update `TestCheckRegistrationUpToDate` (`run_generate_test.go`): the
     "Absent: allowed during the transition" sub-test's assertion flips —
     absence should now produce a problem, not `""`. Rename the sub-test
     comment accordingly (it currently says "allowed during the transition,"
     which is exactly the assumption this step invalidates).

6. **Full verification pass**, same rigor as any other change to this package:
   `gofmt`, `go vet`, `go build ./...`, the full `bigdiffer` suite fresh
   (`go clean -testcache` first, including `TestFullCorpusParity` and
   `TestCheckRegistrationUpToDate`), `-race -short`, `impi` with CI's actual
   flags (`impi --local . --scheme stdThirdPartyLocal --ignore-generated=true
   ./...`), and a live `go run ./internal/tools/bigdiffer -check` /
   `-update` sanity run.

7. **Update `generation-punchlist.md`:** mark item D (and the folded-in items
   13/15) done, then delete the punchlist entirely per its own stated
   lifecycle ("Transient. Delete when the list is empty.") once nothing else
   is outstanding.

8. **Delete this doc.**
