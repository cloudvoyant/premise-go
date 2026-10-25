# premise-connect-grpc-service

A Go service that accepts Connect, gRPC, and gRPC-Web requests.

## Run

```bash
mise run run
```

`ADDR` selects the bind address and defaults to `:8080`. The server supports cleartext HTTP/2 for local gRPC clients. SIGINT starts graceful shutdown.

## Develop

```bash
mise run install
mise run format:check
mise run lint
mise run test
mise run build
```

Formatting and lint tasks use the workspace's root `.golangci.yml`.

## Generate Protocol Code

Run `mise run generate` to regenerate checked-in code from `proto/` with the pinned local plugins. Repeat generation and compare `gen/` for reproducibility. Tests exercise all three protocols through a real HTTP/2 server.

The protobuf `go_package` option stays fixed. `buf.gen.yaml` maps the generated import path through both local plugins. This keeps workspace module substitutions out of encoded protobuf descriptors.

## Workspace

This project is a package in the workspace's root Go module. Keep the same module path when you add projects. Applications live under `apps/`, and libraries live under `libs/`. All packages share dependencies and one repository tag.

`package.public: false` skips project publication tasks. It does not remove source from the module. Use Go `internal/` directories to restrict external imports. Before release, run `GOWORK=off go test ./...` from the workspace root and remove filesystem `replace` directives.
