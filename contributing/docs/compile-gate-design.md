<!-- Copyright IBM Corp. 2021, 2026 -->
<!-- SPDX-License-Identifier: MPL-2.0 -->

# Compile gate design (generation-punchlist.md item 1)

Draft for review — not yet implemented. Deletes into
`suppressed-and-frozen.md`/`new-generation.md` once the design is settled and
built, same as every other punchlist item.

## The problem

Generation is today's gate: `-update` regenerates a candidate for real, and a
generation error (the owned engine returning `err`) is the only failure mode
`decide()` sees. `format.Source` (gofmt-level syntax check) runs inside
generation, but nothing checks that generated code actually **type-checks and
links** against the real module — the SDKs it imports
(`terraform-plugin-framework`, etc.), the shared packages it calls
(`internal/generic`, `internal/identity`, `internal/registry`,
`internal/validators`), or its own package's other files (the single
blank-import `registrations_gen.go` `emit.go` emits, or a naming collision with
a sibling type in the same AWS service package).

`make build` (`go install ./...`) catches this today, but only as a *separate,
manual* step after `-update` has already written the overlay and the tree —
by then the damage (a broken commit) is already made, only caught by the
human running `make build` next. Item 1 is about bigdiffer catching this
itself, in the same pass, before anything is promoted, so a build failure gets
routed through the same never-regress / never-block policy every other
failure mode already gets.

## Why this is not a per-resource gate, unlike generation

Generation and per-artifact validation are naturally per-resource: each
artifact (`resource`/`singular_data_source`/`plural_data_source`) is rendered
independently by the owned engine and either succeeds or fails on its own —
that is why `decide()` can suppress exactly the broken artifact and leave its
siblings alone (item 4).

A compile check cannot honor that granularity, because **Go's unit of build
success/failure is the package, not the file.** Every generated artifact for
one AWS service (e.g. all Logs types) lives in one Go package
(`internal/aws/logs`, `package logs`). If one file in that package fails to
type-check, `go build` fails for the *whole package* — there is no such thing
as "this file's build failed but its 40 siblings' did not." Modeling the
compile gate as one more per-artifact/per-candidate check, parallel to
generation, would be modeling something Go's toolchain does not actually do.

Two consequences follow from that:

1. **Attribution is a parsing problem, not a build-scoping problem.** The
   build either passes or fails for a package; if it fails, `go build`'s
   compiler output (`path/to/file.go:line:col: message`, one line per error)
   is the only way to know *which* file(s) — and therefore which
   candidate(s) — actually broke it. The gate has to run at package (or
   wider) granularity and then attribute failures back down, not the other
   way around.
2. **Running it per-candidate is also a concurrency hazard, not just overkill.**
   Generation is already run concurrently across candidates
   (`generateCorpus`, bounded `errgroup`, proven `-race`-clean — Brick 8). If
   the compile gate also ran per-candidate, two candidates that happen to
   share a service package (routine — most weekly batches touch a handful of
   types in the same handful of services) would each be trying to build a
   package the other is concurrently mutating. That is not a narrow race to
   engineer around; it is fundamentally unsafe for co-package candidates run
   concurrently, because "build this package" is not an operation that
   composes across simultaneous mutations to that same package's files.

## Options considered

| Option | Scope | Race-safe? | Attribution | Cost (measured) |
|---|---|---|---|---|
| A. Per-candidate, concurrent | one candidate's files, run alongside generation | **No** — co-package candidates corrupt each other | Trivial (1:1) | N/A — unsafe, not viable |
| B. Per-candidate, serial (overlay → build → revert, one at a time) | one candidate's package(s) | Yes (serialized) | Trivial (1:1) | ~0.17s/candidate (single package, warm cache) but loses cross-candidate/whole-module interactions |
| C. Per-touched-package, serial, once per batch | every service package touched by the batch, built together | Yes | Requires mapping compiler errors → files → candidates | Not yet measured; between B and D |
| D. Whole-module, serial, once per batch | `go build ./...` | Yes | Requires mapping compiler errors → files → candidates | **~13.7s cold; ~3.1s flat once warm** (measured — the warm figure is `go build`'s module-walk floor, independent of what changed) |

Option A is eliminated outright — it is not a performance tradeoff, it is
broken by construction whenever two candidates share a package, which is the
common case, not an edge case.

Option B avoids the race by giving up concurrency, but at the cost of also
giving up the thing a compile gate is uniquely good for: catching
*cross-file, same-package* interactions (a new type's file colliding with an
existing sibling's identifier, or breaking the package as a whole). Building one
candidate's files alone cannot see the other 40 files in that real package
unless they are present too — so B has to overlay the *whole* package's real
files anyway, undermining the "per-candidate" framing in practice.

Option C scopes the build to just the AWS service packages the batch actually
touched (`go build ./internal/aws/logs/... ./internal/aws/wafv2/...`), which is
almost always a small fraction of the ~200 service packages in the tree.

Option D is the simplest to implement correctly and reason about, is race-free
by the same argument as C (one serial call, not per-candidate), and the measured
cost (~3.1s once warm, run once a week) does not justify the added complexity of
computing and de-duplicating the touched-package set C requires. Its only real
downside versus C is wasted cycles walking untouched packages — but that ~3.1s
is a flat per-invocation floor (module walk, not recompilation), small enough to
absorb for a weekly run.

**Recommendation: D, whole-module, once per batch.** C is a reasonable
alternative if the touched-package computation is judged worth the added
complexity later, but is not proposed as the initial implementation.

## Proposed design (Option D)

The compile gate runs once per `-update` batch, after every candidate has been
staged and before anything real is written. It builds the exact set of `.go`
files the run is about to promote — every staged `*_gen.go` plus the regenerated
`registrations_gen.go` — against the real module graph, and converges on a set
that compiles clean before a single real file is touched. Only when the build is
green does the run promote anything, which is what makes "the promoted tree
compiles" a guarantee rather than an assumption.

### Ordering within `runUpdate`

Everything a successful run produces is staged first, gated, then promoted in
one step:

1. **Candidate loop** (unchanged): each `refreshCandidate` stages its
   successfully-generated artifacts under `stagingDir/out` and the fresh schema
   bytes under `stagingDir/cache`, and yields a per-artifact `gateResult`.
   `decide()` turns each into a tentative `policyDecision`.
2. **Stage the registration file.** Derive the effective row set from the
   current decisions *in memory* (rows with suppressions applied, new rows
   added) and emit `registrations_gen.go` into staging from it. This is the one
   aggregate that affects compilation, so it must be in the gated set — and,
   because dropping an artifact can change the import set (below), it is the one
   aggregate re-emitted on every fixpoint round.
3. **Compile-gate fixpoint** (below): build the staged `.go` set against the
   real tree until it compiles clean, downgrading and dropping any artifact the
   compiler rejects.
4. **Render the overlay text once.** After the fixpoint converges, produce the
   reconciled `all_schemas.hcl` text from the settled decisions
   (`normalizeWithDecisions`). Not compiled, so it sits outside the fixpoint; it
   only has to reflect the final result.
5. **Promote the compiled core once.** With a green build, write the real tree —
   staged code, schema cache, the reconciled overlay, and the gated
   `registrations_gen.go` — together. This is item 12's single atomic promotion.
6. **Emit `import_examples_gen.json` from the promoted state.** Unlike the
   registration file, it is derived from each resource's *schema bytes* (it reads
   the cached schema to recover primary-identifier names), so it must be produced
   after the cache is promoted, from the final overlay + cache — exactly as
   today. It is not compiled and not part of the atomic core promotion; it is a
   deterministic projection that trails it, and is the one output that reads
   promoted schemas rather than in-memory decisions.

   This trailing, non-atomic write is deliberate and acceptable. A crash between
   step 5 and step 6 leaves a correct, compiling, correctly-registered tree with
   a stale `import_examples_gen.json`, which the next `-update`/`-generate`
   regenerates — it is not compiled, so it cannot break `make build` or `-check`,
   and it self-heals rather than lingering as a lasting inconsistency. This is a
   *smaller* window than today, where the overlay and registration file also
   trail `promoteStaged`; the new design pulls those two into the atomic core and
   leaves only this one output outside it. Folding it in too would require a
   staging-aware schema-cache resolver (it reads bytes not yet in the real cache
   at pre-promote time) — not worth closing a sub-second, self-healing window.
   The step-6 write itself should be temp-file-plus-rename so a crash mid-write
   can only leave the file stale-but-valid, never torn (a torn JSON would break
   `-docs`); that is cheap and applies to today's code equally.

A build failure therefore changes *what* gets promoted (fewer artifacts, an
overlay and registration file adjusted to match) but never promotes something
that does not compile, and never blocks the release.

> **Blast radius.** This is a rewrite of the tail of `runUpdate`, not an
> addition. Today that tail writes `all_schemas.hcl`, then re-reads it to derive
> `finalRows`, then emits `registrations_gen.go` and `import_examples_gen.json`
> straight to the real tree. Two things change: (a) `emitRegistration` must stop
> consuming `finalRows` re-read from the *written* overlay and instead take
> in-memory rows + decisions, so it can run per fixpoint round without touching
> disk; (b) the overlay text moves to "produced once after the fixpoint settles,"
> and `import_examples_gen.json` stays *after* promotion (it reads the promoted
> cache), so it is not folded into the gated/atomic set — it trails it.
>
> End-to-end, the per-round work decomposes without new registration plumbing:
> `registrationPackages` already takes `[]resourceRow` and derives the import set
> purely from each row's `suppress_*` flags via `generationPlan` — no schema
> read, no HCL round-trip — so the only new piece is a small in-memory projection
> that applies the current decisions onto the row set (set the `suppress_*`
> flags, add new-type rows) to feed it. That projection is much lighter than the
> full overlay-text render `normalizeWithDecisions` does, and is the only
> per-round decision work; `normalizeWithDecisions` itself runs once.

### How it builds (no synthetic module)

Generated code only type-checks against the real module graph — the
plugin-framework SDK, the `internal/generic`/`identity`/`registry`/`validators`
packages it calls, and its own package's other real files. No standalone temp
module reproduces that graph, so the gate builds in place, reversibly:

- Copy every staged `.go` file over its real destination under `cfg.outputRoot`
  (and the staged `registrations_gen.go` over the real one), recording each
  target's prior bytes — or its prior absence — first.
- `go build ./...` from `cfg.repoRoot`, serially, once per round.
- **Unconditionally revert** via `defer` — restore every recorded file to its
  prior bytes, or delete it if it did not exist — regardless of outcome, so the
  real tree is bit-for-bit unchanged when the gate returns.

The gate is thus the one part of `-update` that touches the real tree
*transiently*; the actual promotion still happens exactly once, afterward. The
only window in which a leftover could survive is a hard process kill (`SIGKILL`)
between overlay and revert — `defer` covers returns and panics, not signals.
Such a leftover surfaces as ordinary uncommitted `git status` output on the next
run, the same pre-existing risk as today's manual `make build`, not silent
corruption.

### The fixpoint: build until clean

`go build`'s unit of success is the package, not the file: one bad file fails
its whole package, and the compiler may stop before reporting every error in it.
A single build-and-attribute pass therefore cannot certify the *remaining*
files — dropping the one file the compiler happened to name can leave other
errors in the same package unreported, and removing a file changes the package
that has to build. So the gate iterates:

1. Build the current staged set (as above).
2. **Green** → done; the staged set is exactly what promotes.
3. **Failure** → parse `go build`'s stderr (`path/to/file.go:line:col: message`,
   one error per line — Go's format is stable and machine-parseable) into the
   set of blamed files, and map each back to the staging path it came from,
   hence to the candidate and the specific artifact (`codeFile` per
   `genArtifact`). Then:
   - For every **attributable** artifact: downgrade its `gateResult` from
     `gateOK` to a new `gateFailedBuild` outcome, **remove its staged file from
     `stagingDir/out`**, recompute its `decide()` (so the overlay will suppress
     exactly that artifact, tagged `reasonBuildFailed`), re-emit the staged
     `registrations_gen.go` from the updated decisions (a newly-added service
     package whose every artifact just failed must drop out of the import set,
     or its blank import would reference an empty package and re-break the
     build), and loop.
   - If **no** blamed file maps to a staged artifact, take the fallback below.

Each non-green round removes at least one artifact from the staged set, so the
loop strictly shrinks and terminates in at most as many rounds as there are
staged artifacts. In practice it converges far faster, because `go build`
reports many errors per pass, not one: measured, `gc` emits up to ~10 errors per
package before it stops that package with `too many errors`, so a round drops a
batch of artifacts, not a single one. That same `too many errors` cutoff is also
why one build can never be assumed to have surfaced *every* error — which is the
second reason the fixpoint exists, not just cross-file masking.

Because the cutoff means a genuinely pathological package (say 40 independently
broken sibling files) still takes several rounds — ~4 at 10 errors/round, each a
full ~3.1s `go build`, ≈13s total — the loop carries a **small round cap**. The
count is already bounded without it (each non-green round removes ≥1 artifact, so
it cannot exceed the staged-artifact count) and the wall-clock is trivial for a
once-a-week run, so the cap is not runaway protection; it is a clean terminator
for the one situation that should never arise — attribution failing to shrink the
set — degrading it to the same conservative exit as an unattributable failure
(below): promote nothing, return a hard error. Promotion is reached only from step 2, so the set that
promotes is always one that built clean — including the regenerated
`registrations_gen.go`, so a new service package is proven both to compile and
to be correctly wired in.

### Attribution fallback

The unattributable case — the compiler blamed a pre-existing base file the batch
never touched, or the generated `registrations_gen.go` itself — is deliberately
conservative: promote nothing, return a hard error. Those are the two situations
that produce it, and both are things to surface, not paper over: a base tree
that was already broken before the run (a regression to investigate — note that
`-check` normalizes the overlay but does *not* build, so the base compiling is
guaranteed only by CI on `main`, not by `-check`), or a bug in bigdiffer's own
`registrations_gen.go` emission (a tool bug). Suppressing every candidate that
merely shares a package with the unattributable error was considered and
rejected as more surprising, not less. This branch is also the loop's backstop
terminator if attribution ever fails to shrink the set.

### Scope: what the gate covers

- **Covered:** every generated package type-checks against the real module
  graph; cross-file, same-package collisions within a service package; and the
  regenerated `registrations_gen.go`.
- **Not covered:** generated `*_gen_test.go` files. `go build ./...` compiles no
  test files by design, matching today's `make build` (`go install ./...`),
  which likewise defers test compilation to `make smoke` / `make test`. A
  generated test that fails to compile is caught there, as it is today. (`go
  vet` would compile tests but is a broader, separate check — `make vet` — not
  part of "does this code build," so it is not used here.)

## What this unblocks

- `build_failed` becomes reachable (`policy.go`'s taxonomy already defines it;
  it has been a dead category since item 3/4 landed).
- `-heal` step 3 (deferred when item 9 shipped): re-probing a suppressed
  artifact can now also run it through the compile gate, not just generation,
  before proposing a `lift`.

## Decisions

- **Whole-module (D), not touched-packages-only (C).** The measured 3.5s
  warm-cache cost, once a week, does not justify computing and de-duplicating
  the touched-package set — work that largely duplicates the error-attribution
  logic the failure path needs regardless. Revisit C only if build time becomes
  a measured problem.
- **`go build ./...`, not `go vet` or `go install`.** Compile-and-discard is the
  right scope for "does it build"; `go install`'s `GOBIN` write buys nothing
  here, and `go vet` is a separate concern (see Scope). Run from `cfg.repoRoot`
  so the whole module graph, including the top-level `internal/provider` package
  that blank-imports every service, is exercised.
- **The fixpoint's inner rounds stay whole-module too, not narrowed to the
  failed packages** — but *not* because the cache makes them free; it doesn't.
  Measured on this repo, `go build ./...` costs a flat ~3.1s whether or not
  anything changed (no-op 3.08s; after touching one file 3.22s), because that
  time is `go build`'s own module walk and staleness check across ~200+
  packages, not recompilation (a single changed package recompiles in ~0.14s).
  So round N+1 costs about the *same* ~3.1s as round 1, not less. Narrowing a
  round to just the changed package(s) genuinely would cut it to ~0.14s — but
  the total is trivial for a once-a-week run even at the pathological cap (~4
  rounds ≈ 13s), so it is not worth the added scoping complexity. And a narrowed
  set could never be *just* the failed package anyway: dropping an artifact
  re-emits `registrations_gen.go`, forcing the top-level `internal/provider`
  package (blank-importing every service) to rebuild regardless.
