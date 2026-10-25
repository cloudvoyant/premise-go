# premise-go Development Guide

## Prerequisites

- [Premise](https://github.com/cloudvoyant/premise) owns generation, contract tests, and release sequencing.
- [Mise](https://mise.jdx.dev/) installs Go and runs the declared tasks.
- Network access supplies Go dependencies and the optional Pulumi SDK tools.

The root Mise configuration pins Premise `v0.6.0`. Each template declares the tools needed after generation. You do not need to build the companion core branch.

## Getting Started

```bash
mise install
mise run build
mise run test
mise run e2e
```

`mise run test` invokes `premise template test` with the pinned CLI. Premise installs each template's tools and dependencies, then runs its contract with `PREMISE_TEMPLATE_TEST=1`. Interactive tasks finish in this mode, and publication tasks do not publish.

`mise run e2e` runs the Go integration suite. It generates all six templates into one module, tests Go and Pulumi consumers, preserves client quality rules, and rejects formatting and complexity violations. It also checks that generated workflows use the stable major action reference.

To test a local core build, set its absolute binary path:

```bash
PREMISE_BIN=/absolute/path/to/premise mise run test
PREMISE_BIN=/absolute/path/to/premise mise run e2e
```

## Project Structure

```text
premise.yaml             # Template declarations and shared client files
mise.toml                # Registry tasks and the released Premise CLI pin
.editorconfig            # Initial workspace editor rules
.golangci.yml            # Initial workspace lint and formatting rules
templates/premise-*/     # Standalone source templates and their tasks
tests/                   # Go integration tests for generated workspaces
.github/workflows/       # Standard Premise commit and merge workflows
docs/                    # Registry development documentation
```

Shared files live at the repository root. The manifest declares `.gitignore` in `workspace_files` and the quality files in `workspace_init_files`. Only the selected template tree becomes an `apps/<name>` or `libs/<name>` project. The registry's root `mise.toml` remains registry-only.

Each source template has a standalone `go.mod` so the common contract can test it in isolation. During generation, the Go backend merges dependencies into the client's root module and removes the project manifests. Native Go install tasks regenerate checksums. The client has one root `go.mod` and `go.sum`, without nested project modules or `go.work`.

The root Mise configuration discovers tasks under `templates/premise-*`, `apps/*`, and `libs/*`. Run `mise tasks --all` to list them.

## Development Workflow

1. Edit shared rules at the repository root or project files under `templates/<name>/`.
2. Run the affected template's formatting, lint, and test tasks.
3. Run the registry checks below.
4. Generate the edited template through an absolute local selector and inspect the client workspace.

```bash
mise run format:check
mise run lint
mise run test
mise run e2e
```

To work on one source template:

```bash
mise run //templates/premise-cobra-cli:format
mise run //templates/premise-cobra-cli:lint
mise run //templates/premise-cobra-cli:test
```

Use `pm generate /absolute/path/to/premise-go:premise-go-lib` to test local changes. Remote selectors use the repository's default branch.

The shared `.golangci.yml` enables gofmt, goimports, static analysis, error checks, HTTP response cleanup, and a complexity limit of 30. Generation copies missing quality files only once. Clients can edit those files, and later projects preserve their contents.

## CI Flows

The standard commit and merge workflows use `cloudvoyant/premise@v0`. Premise maintains that tag after stable core publication. The root Mise hooks run registry checks, and Premise runs the selected template contracts. Checks finish before publication.

Pull requests and unmarked feature pushes validate without publication. A feature push with `[publish-rc]` requests an RC after successful checks. Merges to the default branch use the stable release flow. Generated client workflows also use `@v0` and install the latest stable CLI by default.

## Publishing

All managed Go packages in a generated workspace share one module version and repository tag. External consumers can request one package:

```bash
go get example.com/premise-module/libs/premise-go-lib@v1.2.3
```

The request uses the shared module version. It does not create a separate library release. `package.public` controls project publication tasks and defaults to true. It does not remove source from the module. Questionnaire answers become project metadata in the client's `premise.yaml`.

Before release, run `GOWORK=off go test ./...` from the client root. Do not publish filesystem `replace` directives. Premise publishes Go source through the repository tag and can package application executables as release archives.

The native Pulumi library uses Go mocks and `examples/consumer`. It needs no cloud credentials. The optional provider template also generates Go and Node SDKs. Its Go SDK uses the shared module tag. Node package and plugin distribution require separate release configuration. Keep the Pulumi `pkg/v3` and `sdk/v3` dependency versions aligned across templates.

## Generated Sources

Connect code is checked in under the service template's `gen/` directory. Run `mise run generate` inside that template to regenerate code with its pinned local protobuf and Connect plugins. Repeat generation and compare `gen/` for reproducibility.

Inside the optional Pulumi provider template, run `mise run generate` to rebuild the provider and SDKs. Run `mise run test:sdk` for the Node consumer with Pulumi mocks. The normal `test` task includes that consumer. Neither consumer deploys cloud resources. See the [provider README](../templates/premise-pulumi-component-package/README.md) for shared-version SDK generation and distribution requirements.
