# premise-htmx-app

Each project is its own Go module. Choose a canonical module path under your repository and a publication policy. `private-local` modules stay in the workspace; `public` modules receive directory-prefixed tags at the shared workspace version.

Run `mise run install`, `mise run test`, and `mise run build`. The workspace's `go.work` resolves local modules. External consumers run `go get example.com/premise-module@v1.2.3`. Verify public modules with `GOWORK=off go test ./...` before release. Do not publish filesystem replace directives.

`ADDR` selects the HTTP bind address (default `:8080`). Templates and htmx 2.0.8 are embedded in the binary, so deployment needs only the executable. SIGINT starts graceful shutdown. `POST /greet` returns an HTML fragment with escaped input. The vendored htmx script is distributed under its upstream 0BSD license.
