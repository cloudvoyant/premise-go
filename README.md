# premise-go

A Go template registry for [Premise](https://github.com/cloudvoyant/premise). Generate applications and libraries under `apps/<name>` and `libs/<name>` with shared dependencies and one repository version.

## Templates

| Template                           | Kind | Description                                               |
| ---------------------------------- | ---- | --------------------------------------------------------- |
| `premise-cobra-cli`                | app  | Cobra command-line application.                           |
| `premise-connect-grpc-service`     | app  | HTTP service for Connect, gRPC, and gRPC-Web.             |
| `premise-htmx-app`                 | app  | Web application with embedded HTML templates and htmx.    |
| `premise-go-lib`                   | lib  | Go library for external or internal use.                  |
| `premise-pulumi-component-lib`     | lib  | Go Pulumi components with a Go consumer and mock tests.   |
| `premise-pulumi-component-package` | lib  | Optional Pulumi provider with generated Go and Node SDKs. |

Generated projects share one root `go.mod`, `go.sum`, dependency set, and release tag. Each project has its own Mise tasks. External clients can import one library package at the shared module version. Generated projects do not need `go.work` or separate library tags.

Shared files live at the repository root, as they do in the Bun and Cargo registries. `workspace_init_files` declares `.editorconfig` and `.golangci.yml` as initial client defaults. Generation copies missing files, preserves existing client rules, and continues to merge dependencies on later additions.

## Requirements

- Premise `v0.6.0` or later on `PATH` as `pm`.
- [Mise](https://mise.jdx.dev/) to install the declared toolchains.
- Network access to download Go dependencies and the optional Pulumi SDK tools.

## Generate a Project

```bash
pm generate cloudvoyant/premise-go
```

Select a template and enter a project name and workspace module path. Use the same module path for every Go project in that workspace. To select a template directly:

```bash
pm generate cloudvoyant/premise-go:premise-go-lib
```

Choose `premise-pulumi-component-lib` for components consumed by other Go Pulumi projects. Choose the optional `premise-pulumi-component-package` only when you need generated SDKs for other languages.

Libraries ask whether to run publication tasks. `package.public: false` skips those tasks, but the source remains part of the shared module. Use Go `internal/` directories to restrict external imports.

## Development

The registry pins the released Premise CLI `v0.6.0` through Mise. CI uses the maintained `cloudvoyant/premise@v0` action tag. See the [development guide](docs/development-guide.md) for installation, validation, workspace rules, and publication.

```bash
mise install
mise run test
mise run e2e
```

The root test uses the same `premise template test` contract as Bun and Cargo. The Go integration suite also generates all six templates into one workspace and tests the shared module and quality rules.

## License

[MIT](LICENSE)
