<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# bigdiffer runbooks

Task-oriented runbooks for `bigdiffer` (`internal/tools/bigdiffer`), the tool
that generates the Terraform AWS Cloud Control Provider from the live
CloudFormation registry. For what bigdiffer is and the available commands,
see [the tool's README](../../internal/tools/bigdiffer/README.md); for why
it's built this way, see [bigdiffer-design.md](bigdiffer-design.md).

Pick the runbook for the job you're doing — they are independent, not
sequential steps of one bigger process:

<!--mdtoc: begin-->
* [Weekly release prep](#weekly-release-prep)
* [Adding an example or changing a doc template](#adding-an-example-or-changing-a-doc-template)
* [Changing the codegen](#changing-the-codegen)
* [Testing after a Go dependency or version bump](#testing-after-a-go-dependency-or-version-bump)
* [Checking the overlay's bookkeeping](#checking-the-overlays-bookkeeping)
* [Revisiting an old suppression or freeze](#revisiting-an-old-suppression-or-freeze)
* [Fallback: the legacy process](#fallback-the-legacy-process)
* [Reading the report](#reading-the-report)
* [How it works](#how-it-works)
<!--mdtoc: end-->

> [!TIP]
> `make bigdiffer-sync`, `make bigdiffer-reconcile`, `make bigdiffer-check`,
> `make bigdiffer-docs`, `make bigdiffer-lint`, and `make bigdiffer-recheck`
> are shortcuts for the `go run ./internal/tools/bigdiffer …` commands used
> below — prefer them over the bare `go run` form: `make` pins `GOTOOLCHAIN`
> to `go.mod`'s version for every recipe (see `GNUmakefile`), so generated Go
> source is formatted identically to CI regardless of your locally installed
> Go. A flag combination with no dedicated target (`-no-docs`,
> `-check-docs-full`, `-recheck-all`) still needs the bare `go run` form,
> shown where used below; `-sync`/`-reconcile`/`-check` verify the toolchain
> themselves either way (`goversion.go`), so a mismatch is refused rather
> than silently producing CI-diverging output, but matching `make`'s pin
> from the start avoids hitting that refusal at all.

## Weekly release prep

The job: bring the provider up to date with whatever AWS added or changed
this week, and get a PR ready.

### 1. Setup

```sh
git switch main
git pull
make tools
make newbranch
export AWS_DEFAULT_REGION=us-east-1   # bigdiffer always queries us-east-1
```

Ensure valid AWS credentials are set in the environment.

### 2. Sync

```sh
make bigdiffer-sync
```

This refreshes the schema cache, tracks new CloudFormation types, gates and
regenerates the **whole** corpus (not just changed types — every type is
generate-and-compile-checked every run, so a machinery change surfaces
immediately instead of aging silently), rewrites `all_schemas.hcl`, and
regenerates documentation (import-example docs, `terraform fmt`,
`tfplugindocs generate`) — all in one pass, taking roughly 13 minutes
(dominated by the AWS discovery crawl; a progress bar shows it). Pass
`-no-docs` (there's no dedicated target for a flag, only for a command, so
use `go run ./internal/tools/bigdiffer -sync -no-docs` directly, or
`GOTOOLCHAIN=go$(cat .go-version) go run ./internal/tools/bigdiffer -sync
-no-docs` to pin the toolchain the way `make` would) to
skip the documentation step for a faster inner loop; run
[`-docs`](#adding-an-example-or-changing-a-doc-template) separately before
committing in that case. It prints a [report](#reading-the-report) of what
changed and what, if anything, was frozen or suppressed.

A schema-changed type that fails generation/compile freezes or suppresses just
that type, same as before. A **schema-unchanged** type that fails means the
generation machinery itself regressed — a template, `codegen/`, or a Go
toolchain change — so `-sync` aborts the **entire run** and promotes nothing at
all, not even the types that compiled cleanly this run, and reports exactly
which types and artifacts failed. That is not a bug to work around: an engine
regression makes every type it touched this run untrustworthy, not just the
one that happened to fail to compile. Fix or revert the machinery change (see
[Changing the codegen](#changing-the-codegen)), then re-run.

A documentation failure is different in kind: it happens *after* everything
above has already promoted successfully, and bigdiffer never commits on its
own, so a docs failure is a loud, recoverable error over an otherwise-valid,
uncommitted working tree, not a broken commit. Fix the underlying problem
(a missing tool, a template bug), then re-run `make bigdiffer-docs` alone to
finish — no need to redo the sync.

Review the report together with `git diff`. For each frozen/suppressed type,
open a GitHub issue with the reason (see
[example issue #2070](https://github.com/hashicorp/terraform-provider-awscc/issues/2070));
`internal/update/suppressions_checkout.txt` pins are still honored.

Files this changes:

* `internal/service/cloudformation/schemas/*.json` — refreshed schema cache
* `internal/provider/all_schemas.hcl` — new blocks + policy annotations
* `internal/aws/**/*_gen.go` — regenerated code for every type the gate
  promotes this run (new/changed types, plus any type whose output changed
  from a machinery fix — schema-unchanged types otherwise regenerate
  byte-identical and are a no-op)
* `internal/provider/import_examples_gen.json` — import-examples aggregate
* `internal/provider/registrations_gen.go` — the blank-import registration file
* `CHANGELOG.md` — a drafted `FEATURES:` block for newly user-visible artifacts
* `examples/**`, `docs/**` — regenerated documentation (unless `-no-docs`)

> [!NOTE]
> `registrations_gen.go` supersedes the legacy
> `internal/provider/{resources,singular_data_sources,plural_data_sources}.go`
> directive files, which bigdiffer does **not** touch. Both may be present
> during the transition — duplicate blank imports are legal Go. Commit
> `registrations_gen.go`. `bigdiffer -lint` (run in CI) re-emits this file and
> fails if the committed copy is stale.

### 3. Build and smoke test

```sh
make build
make smoke
```

`-sync` already compiled the generated code before writing it, so `make
build` here is a confirmation, and `make smoke` adds test compilation and
acceptance smoke tests. If `make build` fails here, investigate rather than
hand-suppressing — `-sync` would already have suppressed a genuinely
non-compiling type, or aborted entirely on a machinery regression.

### 4. CHANGELOG and version

`-sync` already drafted the `FEATURES:` bullets into `CHANGELOG.md`'s top,
in-progress version block (e.g. `## 1.100.0 (Unreleased)`). By hand:

1. Add the PR number link
   (`([#1234](https://github.com/hashicorp/terraform-provider-awscc/pull/1234))`)
   once the PR exists.
2. Add any `NOTES:` or breaking-change entries.
3. Update `version/VERSION` to match.

If this run promoted nothing newly user-visible, no `FEATURES:` section is
added.

Running `-sync` again within the same still-open cycle (no release cut in
between) is safe: the new run's bullets merge into the top block's
`FEATURES:` section's existing contiguous run of recognized bullets —
deduplicated, re-sorted — rather than requiring a release between every
sync. A `provider: ...` note ahead of the bullets (the common shape
throughout `CHANGELOG.md`'s history) is left exactly where it was. Only if
the recognized bullets are split into more than one group by unrecognized
content in between does `-sync` still error rather than guess which group to
merge into; add the new bullets by hand in that case.

### 5. Commit and open a pull request

Review `git status` and `git diff`, then commit — either as a single commit or
grouped logically, for example:

```sh
git add internal/service/cloudformation/schemas internal/provider/all_schemas.hcl
git commit -m "Refresh CloudFormation schemas"

git add internal/aws internal/provider/registrations_gen.go internal/provider/import_examples_gen.json
git commit -m "Regenerate resources and data sources"

git add examples docs CHANGELOG.md version/VERSION
git commit -m "Regenerate documentation and update changelog"
```

Open a pull request and verify CI passes; once merged, cut the release.

## Adding an example or changing a doc template

The job: an examples-only or doc-template change, with no code regeneration
at all — regenerate documentation from what's already committed.

Requires Terraform `v1.14`+ and `tfplugindocs` on `PATH` (both installed by
`make tools`):

```sh
make bigdiffer-docs
```

This regenerates import-example docs, runs `terraform fmt`, and runs
`tfplugindocs generate` — the same documentation step `-sync`/`-reconcile`
run automatically as their last step, run here on its own. Use this runbook
when nothing about the generated *code* needs to change, or to finish a
`-sync`/`-reconcile` run that failed only at the documentation step. Changes
`examples/**` and `docs/**`.

## Changing the codegen

The job: you changed a template, `codegen/`, or the naming package, and want
to see what it would break before landing it, then land it.

First, preview the change without writing anything:

```sh
make bigdiffer-check
```

This runs the same generate → compile → decide gate `-sync`/`-reconcile` use,
offline, with promotion switched off. It reports every failure, then exits
non-zero on either a failure **or** any regenerated output — including
`internal/provider/import_examples_gen.json` — that differs from what's
committed, even if it compiles: the "you changed the engine but haven't
`-reconcile`'d and committed yet" case. A non-engine change touches nothing
`-check` would regenerate differently, so it passes with a clean diff.

If the change might affect the rendered registry docs themselves — a
template, a schema-description path, anything `tfplugindocs` reads — add
`-check-docs-full` to also render `docs/` (via a real built binary) and diff
against committed (no dedicated target for a flag, so use `go run` directly,
or pin the toolchain yourself the way `make` would):

```sh
GOTOOLCHAIN=go$(cat .go-version) go run ./internal/tools/bigdiffer -check -check-docs-full
```

This is markedly heavier than plain `-check` (a real whole-provider build
plus a schema-extraction call, a couple of minutes rather than under a
minute), so reach for it when the change plausibly touches docs rendering,
not as a routine step every time.

Once `-check` finds a real diff (or you already know you changed something),
land it:

```sh
make bigdiffer-reconcile
```

This runs every committed type through the identical gate, offline, and
promotes the result if the whole run is clean. Unlike the old, deleted
`-generate` mode, `-reconcile` **does** compile-gate: any failure — there is
no schema change to distinguish here, so every failure is a machinery
regression — aborts the entire run and promotes nothing, with the same
per-type-and-artifact blame `-sync` reports. It does not query AWS and writes
the same files `-sync` would (schema cache, generated code, registration
file, import-examples aggregate, `all_schemas.hcl`, documentation, and a
CHANGELOG draft for any backlog artifact the fix newly recovers), just from
the committed base instead of a live crawl. Repeat `-check` until it is
clean, then commit.

`-sync` works too if you also want to pick up whatever changed live in AWS
this week in the same pass — see [Weekly release prep](#weekly-release-prep).

## Testing after a Go dependency or version bump

The job: confirm a `go.mod`/`go.sum` dependency change, or a bump to the Go
version pinned in `go.mod`/`tools/go.mod`/`.go-version`, didn't alter
generated output or break the build.

If you already updated the pin (the usual order — bump the version files
first, matching whatever `terraform-provider-aws` or upstream Go itself just
shipped, then verify), `make` will pick it up automatically:

```sh
make bigdiffer-check
```

If you're instead trial-running a *prospective* bump before touching the
pin, run under that candidate version explicitly so `make`'s own
`GOTOOLCHAIN` pin (still the old, committed version) doesn't mask the very
thing you're testing:

```sh
GOTOOLCHAIN=go1.27.0 go run ./internal/tools/bigdiffer -check
```

A clean pass with a zero diff means the bump changed nothing bigdiffer can
see. If it reports a diff or a failure, treat it exactly like
[Changing the codegen](#changing-the-codegen): a dependency/toolchain bump is a
machinery change from bigdiffer's point of view, whether or not any of your
own template/codegen code moved. Run `make bigdiffer-reconcile` (after
updating the pin, if you were trial-running one) to land the (hopefully
inert) regeneration, then commit — a real diff here is worth understanding
before committing, since it means the bump changed what the provider emits.

## Checking the overlay's bookkeeping

The job: not "would regenerating break anything," but "is
`all_schemas.hcl` itself consistent, sensible, and explained right now" —
no live AWS check, no regeneration at all.

```sh
make bigdiffer-lint
```

This re-sorts and re-formats the overlay against itself (no rows are added,
since there is no discovery here) and fails if the committed file is not
already exactly that normalized shape, if `registrations_gen.go` is stale
relative to the overlay, or if any suppressed/frozen fact has no reason
recorded. It also reports (advisory, not a failure) any row retained in the
overlay with nothing explaining why it's still there. Offline, writes
nothing — this is the check CI runs on every PR.

If it fails on formatting/registration staleness, running `-sync` or
`-reconcile` re-normalizes and re-emits both as a side effect of any real
run; there is no separate "just fix the formatting" command, since
`-lint`'s whole point is that the overlay should already be in the shape
those commands leave it in. If it fails on a reason-less suppression/freeze,
see [Revisiting an old suppression or freeze](#revisiting-an-old-suppression-or-freeze)
below.

## Revisiting an old suppression or freeze

The job: something was marked broken, excluded, or frozen a while back —
find out whether that's still true, or has quietly become fine, without
guessing or trusting stale memory.

```sh
make bigdiffer-recheck
```

By default this re-probes every suppressed/frozen fact that has **no**
recorded reason (or one tagged `unknown`): a structural check first (does
the schema shape still justify the suppression), then a real, isolated
regeneration-and-compile attempt against the committed schema cache, then —
for whatever's left — migrating any existing free-form comment into a
proper reason rather than guessing one. Every result is a **proposal**,
printed to the report; nothing is ever written to `all_schemas.hcl`
automatically. Offline except reading the already-committed schema cache —
no AWS, no live discovery.

To revisit a decision that already has a real, specific reason recorded —
not just the reason-less backlog — widen the scope (no dedicated target for
a flag, so use `go run` directly, or pin the toolchain yourself the way
`make` would):

```sh
GOTOOLCHAIN=go$(cat .go-version) go run ./internal/tools/bigdiffer -recheck -recheck-all
```

Useful after a schema-generating engine change, a suppression whose
underlying AWS issue you believe was fixed, a periodic audit of old
freezes, or after any `-sync`/`-reconcile` run reports a new
`generation_failed` suppression on a type with more than one artifact (see
bigdiffer-design.md's "Isolation gap, fixed" for why a multi-artifact type
can need this). `-recheck-all` re-probes everything active, including facts a
human already explained, on the same "propose, never guess-and-write"
terms.

Review the proposals, then apply the ones you accept by hand-editing
`all_schemas.hcl` (clearing `suppress_*_generation`/`frozen_since`, or just
updating the `*_reason`/`frozen_reason` text) — `-recheck` never edits the
file itself. See
[suppressed-and-frozen.md](suppressed-and-frozen.md#-recheck-re-probe-and-fill-gaps)
for the full reason taxonomy and exactly what each proposal category means.

## Fallback: the legacy process

The job: bigdiffer is unavailable or misbehaving, and the weekly release
still needs to ship.

The legacy machinery is intact and unaffected by bigdiffer; it produces the
same result. Follow [generating-the-provider.md](generating-the-provider.md)
instead (`make cleanschemas suppressions schemas`, `make biglister`,
`make resources`, `make singular-data-sources plural-data-sources`,
`make docs-all`, …).

Interoperating between the two:

* The **schema cache** (`internal/service/cloudformation/schemas/*.json`) and the
  **overlay** (`all_schemas.hcl`) use the identical on-disk format, so a cache and
  overlay produced by bigdiffer are consumed unchanged by the legacy `make`
  targets, and vice versa. No conversion is needed.
* Before committing a legacy-only run, delete bigdiffer's registration file so
  registration comes solely from the regenerated legacy directive files:

  ```sh
  rm -f internal/provider/registrations_gen.go
  ```

  `make resources` regenerates `internal/provider/resources.go` (and the
  data-source equivalents) complete with every current type, so the legacy path is
  self-contained.

> [!NOTE]
> The legacy `make bigdiffer` target is an unrelated helper that `diff`s dated
> `available_schemas.<date>.hcl` snapshots — it is **not** the
> `internal/tools/bigdiffer` tool described here.

## Reading the report

`-sync`/`-reconcile` narrate each phase (`==>` lines) and print a
reconciliation report to stderr:

* **Detected changes — new / changed / unchanged / frozen / missing.** Every
  type is generate-and-compile-gated every run; New and Changed also refresh
  their schema bytes from AWS, Unchanged only re-verifies the machinery still
  produces the same output, Frozen (pinned via `frozen_since`) are never
  re-evaluated at all, and Missing means discovery could not describe the type.
* **Per-type lines** for anything frozen or suppressed, with the policy summary
  and reason — these are your issue-filing worklist.
* **Anomalies** — retained-but-unexplained rows (in the overlay, gone from AWS,
  with no `frozen_since`/`non_provisionable`/checkout pin), duplicate blocks, and
  naming-invariant violations. Investigate these before merging.

## How it works

Briefly: bigdiffer reconciles the live CloudFormation registry against
`internal/provider/all_schemas.hcl` in memory, gates the whole corpus
(generate and compile) every run, and lets one policy decide what to freeze
or suppress — a broken *schema* never blocks a release, but a regression in
the generation *machinery itself* fails the whole run rather than shipping
output from a generator just caught being broken. Documentation is part of
the same gated pipeline, not a separate step to remember: `-sync`/`-reconcile`
render it automatically after promoting, and `-check` enforces that a
committed `import_examples_gen.json` still matches what the current engine
would produce.

For the full design — why it's built this way, what it replaces, and what's
still open — see [bigdiffer-design.md](bigdiffer-design.md).
