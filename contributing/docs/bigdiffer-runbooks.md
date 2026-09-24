<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer runbooks

Task-oriented runbooks for `bigdiffer` (`internal/tools/bigdiffer`), the tool that generates the Terraform AWS Cloud Control Provider from the live CloudFormation registry. For what bigdiffer is and the available commands, see [the tool's README](../../internal/tools/bigdiffer/README.md); for why it's built this way, see [bigdiffer-design.md](bigdiffer-design.md).

Pick the runbook for the job you're doing — they are independent, not sequential steps of one bigger process:

<!--mdtoc: begin-->
* [Weekly release prep](#weekly-release-prep)
* [Adding an example or changing a doc template](#adding-an-example-or-changing-a-doc-template)
* [Changing the codegen](#changing-the-codegen)
* [Testing after a Go dependency or version bump](#testing-after-a-go-dependency-or-version-bump)
* [Checking the overlay's bookkeeping](#checking-the-overlays-bookkeeping)
* [Revisiting an old suppression or freeze](#revisiting-an-old-suppression-or-freeze)
* [Fallback: the legacy process](#fallback-the-legacy-process)
* [Reading the report](#reading-the-report)
<!--mdtoc: end-->

> [!TIP]
> Prefer the `make bigdiffer-*` targets shown below: they pin `GOTOOLCHAIN` to the `go.mod` Go version, so generated code matches CI. For a flag with no target, use bare `go run ./internal/tools/bigdiffer …` (shown where needed) and pin the toolchain yourself with `GOTOOLCHAIN=go$(cat .go-version)` — a different local Go produces false diffs.
>
> Whole-corpus commands (`-check`, `-reconcile`) and the test suite launch ~1600 short-lived subprocesses; if endpoint security interrupts a local run, rely on CI's `bigdiffer` workflow for full-corpus verification.

## Weekly release prep

The job: bring the provider up to date with whatever AWS added or changed this week, and open a PR.

### 1. Setup

```sh
git switch main
git pull
make tools
make newbranch
export AWS_DEFAULT_REGION=us-east-1   # bigdiffer always queries us-east-1
```

Ensure valid AWS credentials are set in the environment.

### 2. Sync and regenerate

```sh
make bigdiffer-sync
```

Refreshes the AWS schemas, regenerates the provider and documentation, and prints a report. Expect about 13 minutes.

Review the report and `git diff`. If the sync itself fails, fix the reported generation problem and rerun — do not work around it. For each frozen or suppressed type, open a GitHub issue with the reported reason.

For a faster inner loop, skip documentation with `go run ./internal/tools/bigdiffer -sync -no-docs` (no target for the flag; pin the toolchain as above), then run `make bigdiffer-docs` before committing.

### 3. Build and smoke test

```sh
make build
make smoke
```

`-sync` already compiled the generated code, so `make build` is a confirmation and `make smoke` adds test compilation and acceptance smoke tests. If `make build` fails here, investigate rather than hand-suppressing.

### 4. CHANGELOG and version

`-sync` drafts the `FEATURES:` bullets into `CHANGELOG.md`'s top, in-progress version block. By hand: add any `NOTES:` or breaking-change entries, and update `version/VERSION` to match. If the run promoted nothing user-visible, no `FEATURES:` section is added.

Re-running `-sync` within the same open cycle is safe: new bullets merge into the top block's existing `FEATURES:` run, deduplicated and re-sorted. If unrecognized content splits the bullets into more than one group, `-sync` errors instead of guessing — add the new bullets by hand in that case.

### 5. Commit and open a pull request

```sh
make bigdiffer-commit
```

Commits the run's output as reviewable, path-grouped commits (schemas, resources, data sources, docs), skipping empty groups and refusing to run on `main`. Then review `git status`/`git diff`, push, and open a PR; once CI passes and it merges, cut the release.

**End of runbook.**

---

## Adding an example or changing a doc template

The job: an examples-only or doc-template change, with no code regeneration — regenerate documentation from what's already committed.

Requires Terraform `v1.14`+ and `tfplugindocs` on `PATH` (both from `make tools`):

```sh
make bigdiffer-docs
```

Regenerates import-example docs, runs `terraform fmt`, and runs `tfplugindocs generate` — the same documentation step `-sync`/`-reconcile` run automatically, run here on its own. Use it when no generated *code* needs to change, or to finish a `-sync`/`-reconcile` that failed only at the documentation step. Changes `examples/**` and `docs/**`.

**End of runbook.**

---

## Changing the codegen

The job: you changed a template, `codegen/`, or the naming package — see what it breaks before landing it, then land it.

Preview offline, writing nothing:

```sh
make bigdiffer-check
```

Runs the same generate → compile → decide gate as `-sync`/`-reconcile`, promotion off, and exits non-zero on any failure or on any regenerated output (including `import_examples_gen.json`) that differs from committed. A non-engine change passes with a clean diff.

If the change might affect the rendered registry docs (a template, a schema-description path — anything `tfplugindocs` reads), also render and diff `docs/`. This is markedly heavier (a real whole-provider build, a couple of minutes), so reach for it only when docs rendering is plausibly affected:

```sh
GOTOOLCHAIN=go$(cat .go-version) go run ./internal/tools/bigdiffer -check -check-docs-full
```

Then land it:

```sh
make bigdiffer-reconcile
```

Runs every committed type through the identical gate, offline, and promotes if the whole run is clean; any failure is a machinery regression that aborts the run and promotes nothing. Writes the same files `-sync` would (schema cache, generated code, registration file, import-examples aggregate, `all_schemas.hcl`, documentation, and a CHANGELOG draft for any backlog artifact the fix recovers), from the committed base rather than a live crawl. Repeat `-check` until clean, then `make bigdiffer-commit`.

To also pick up this week's live AWS changes in the same pass, use `-sync` instead — see [Weekly release prep](#weekly-release-prep).

**End of runbook.**

---

## Testing after a Go dependency or version bump

The job: confirm a `go.mod`/`go.sum` change, or a bump to the pinned Go version (`go.mod`/`tools/go.mod`/`.go-version`), didn't alter generated output or break the build.

If you already updated the pin (the usual order), `make` picks it up:

```sh
make bigdiffer-check
```

To trial a *prospective* bump before touching the pin, run under the candidate version explicitly so `make`'s pin doesn't mask the very thing you're testing:

```sh
GOTOOLCHAIN=go1.27.0 go run ./internal/tools/bigdiffer -check
```

A clean, zero-diff pass means the bump changed nothing bigdiffer emits. A diff or failure is a machinery change — treat it like [Changing the codegen](#changing-the-codegen): run `make bigdiffer-reconcile` (after updating the pin, if you were trial-running one) to land the regeneration, then commit. A real diff here is worth understanding before committing, since it means the bump changed what the provider emits.

**End of runbook.**

---

## Checking the overlay's bookkeeping

The job: not "would regenerating break anything," but "is `all_schemas.hcl` itself consistent and fully explained right now" — no AWS, no regeneration.

```sh
make bigdiffer-lint
```

Re-sorts and re-formats the overlay against itself and fails if the committed file isn't already that shape, if `registrations_gen.go` is stale relative to the overlay, or if any suppressed/frozen fact has no reason recorded. It also flags (advisory, not a failure) any retained row with nothing explaining why it's still there. Offline, writes nothing — CI runs this on every PR.

A formatting or registration-staleness failure is fixed as a side effect of any real `-sync`/`-reconcile` run; there is no separate "just fix the formatting" command. For a reason-less suppression or freeze, see [Revisiting an old suppression or freeze](#revisiting-an-old-suppression-or-freeze).

**End of runbook.**

---

## Revisiting an old suppression or freeze

The job: something was marked broken, excluded, or frozen a while back — find out whether that's still true, without guessing.

```sh
make bigdiffer-recheck
```

Re-probes every suppressed/frozen fact that has no recorded reason (or one tagged `unknown`): a structural check first, then a real isolated regeneration-and-compile against the committed schema cache, then migrating any existing free-form comment into a proper reason. Every result is a *proposal* printed to the report; nothing is written to `all_schemas.hcl` automatically. Offline except reading the committed cache.

To also revisit facts that already have a real, specific reason, widen the scope (no target for the flag):

```sh
GOTOOLCHAIN=go$(cat .go-version) go run ./internal/tools/bigdiffer -recheck -recheck-all
```

Useful after an engine change, a suppression whose upstream AWS issue you believe was fixed, a periodic audit of old freezes, or when a `-sync`/`-reconcile` reports a new `generation_failed` suppression on a multi-artifact type.

Review the proposals and apply the ones you accept by hand-editing `all_schemas.hcl` (clearing `suppress_*_generation`/`frozen_since`, or updating the `*_reason`/`frozen_reason` text) — `-recheck` never edits the file itself. See [suppressed-and-frozen.md](suppressed-and-frozen.md#-recheck-re-probe-and-fill-gaps) for the full reason taxonomy.

**End of runbook.**

---

## Fallback: the legacy process

The job: bigdiffer is unavailable or misbehaving, and the weekly release still needs to ship.

The legacy machinery is intact and unaffected by bigdiffer, and produces the same result. Follow [generating-the-provider.md](generating-the-provider.md) instead (`make cleanschemas suppressions schemas`, `make biglister`, `make resources`, `make singular-data-sources plural-data-sources`, `make docs-all`, …), and commit with the legacy `make commit*` targets.

The two paths interoperate:

* The schema cache (`internal/service/cloudformation/schemas/*.json`) and the overlay (`all_schemas.hcl`) use the identical on-disk format, so artifacts produced by one are consumed unchanged by the other. No conversion is needed.
* Before committing a legacy-only run, delete bigdiffer's registration file so registration comes solely from the regenerated legacy directive files (both may coexist during the transition):

  ```sh
  rm -f internal/provider/registrations_gen.go
  ```

  `make resources` regenerates `internal/provider/resources.go` (and the data-source equivalents) with every current type, so the legacy path is self-contained.

When bigdiffer fully replaces the legacy process, retire the legacy generator targets and the `make commit*` targets along with it.

> [!NOTE]
> The legacy `make bigdiffer` target is an unrelated helper that `diff`s dated `available_schemas.<date>.hcl` snapshots — not the `internal/tools/bigdiffer` tool described here.

**End of runbook.**

---

## Reading the report

`-sync`/`-reconcile` narrate each phase (`==>` lines) and print a reconciliation report:

* **Detected changes — new / changed / unchanged / frozen / missing.** Every type is generate-and-compile-gated every run; New and Changed also refresh their schema bytes from AWS, Unchanged only re-verifies the machinery still produces the same output, Frozen (pinned via `frozen_since`) are never re-evaluated, and Missing means discovery could not describe the type.
* **Per-type lines** for anything frozen or suppressed, with the policy summary and reason — your issue-filing worklist. `-sync` also repeats these as an action-item recap at the end of the run.
* **Anomalies** — retained-but-unexplained rows (in the overlay, gone from AWS, with no `frozen_since`/`non_provisionable`/checkout pin), duplicate blocks, and naming-invariant violations. Investigate these before merging.

Full output — including the external tools' per-file lines — is written to a gitignored `.bigdiffer-<command>.log`; the console shows only the structural lines above.
