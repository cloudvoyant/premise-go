# premise-htmx-app

A Go web application with embedded HTML templates and htmx.

## Run

```bash
mise run run
```

Open `http://localhost:8080`. `ADDR` changes the bind address. `POST /greet` returns an HTML fragment with escaped input. SIGINT starts graceful shutdown.

The executable embeds its HTML templates and htmx script. Deployment needs only the binary. The vendored htmx script uses its upstream 0BSD license.

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
