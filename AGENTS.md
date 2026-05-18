# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.bytebuilders.dev/uibuilder-tools` — a developer **CLI** that supports the `ui-wizards` / `ui-builder` ecosystem. The single binary `uibuilder-tools` provides subcommands to lint and validate editor charts (cross-checking `ui/create-ui.yaml`, `ui/edit-ui.yaml`, etc. against `values.openapiv3_schema.yaml`), debug JSON-schema input, and bump versions.

## Architecture

Flat root-level Cobra CLI — no `cmd/` or `pkg/` split.

- `main.go` — entry point.
- `root.go` — Cobra root, wires subcommands: `check`, `debug`, `update-version`, and `version`.
- `check.go` — validates editor charts: walks `wizardDir`, compares each chart's `ui/*.yaml` (`create-ui.yaml`, `edit-ui.yaml`, `language.yaml`) against `values.openapiv3_schema.yaml`, optionally resolves `$ref`s. Used in pre-commit / CI for `ui-wizards`.
- `debug.go` — pretty-prints an embedded UI-schema snippet; scratch space for debugging JSON-schema work.
- `update_version.go` — bulk version bumps across chart files.
- `hack/`, `Makefile` — standard AppsCode build harness.

## Common commands

- `make build` / `make all-build` — host or all-platform build.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make verify` — module-tidy verification.
- `make ci` — full CI pipeline.
- Run locally: `go run . check --wizard-dir=../ui-wizards/charts` (or pass `--ui-file` / `--schema-file` for a single chart).

## Conventions

- Module path is `go.bytebuilders.dev/uibuilder-tools` (vanity URL); imports must use that.
- Single-binary, single-package (`package main`) at the root — keep new subcommands as sibling files under root, not in `cmd/` subdirectories.
- Primary consumer is `ui-wizards`'s CI — if you change the `check` flag surface or exit codes, sync with that repo's hooks.
- License: `LICENSE`. Sign off commits (`git commit -s`); contributions follow the DCO (`DCO`, `CONTRIBUTING.md`).
