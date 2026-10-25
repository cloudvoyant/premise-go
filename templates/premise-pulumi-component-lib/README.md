# premise-pulumi-component-lib

A native Go Pulumi component library with a Go consumer and mock tests.

## Use

From another Go Pulumi project, fetch a released package:

```bash
go get example.com/premise-module/libs/premise-pulumi-component-lib@v1.2.3
```

Import the package as `component` and call `component.NewGreeting`. See `examples/consumer` for a complete Go consumer. Mock tests do not deploy cloud resources or require cloud credentials.

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
