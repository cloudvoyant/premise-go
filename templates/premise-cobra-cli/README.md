# premise-cobra-cli

A minimal Go command-line application built with [Cobra](https://github.com/spf13/cobra).

## Run

```bash
mise run run
```

Install a released application with `go install example.com/premise-module/apps/premise-cobra-cli@v1.2.3`.

## Develop

```bash
mise run install
mise run format:check
mise run lint
mise run test
mise run build
```

Formatting and lint tasks use the workspace's root `.golangci.yml`.

## Workspace

This project is a package in the workspace's root Go module. Keep the same module path when you add projects. Applications live under `apps/`, and libraries live under `libs/`. All packages share dependencies and one repository tag.

`package.public: false` skips project publication tasks. It does not remove source from the module. Use Go `internal/` directories to restrict external imports. Before release, run `GOWORK=off go test ./...` from the workspace root and remove filesystem `replace` directives.
