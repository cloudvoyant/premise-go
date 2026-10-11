# premise-connect-grpc-service

Each project is its own Go module. Choose a canonical module path under your repository and a publication policy. `private-local` modules stay in the workspace; `public` modules receive directory-prefixed tags at the shared workspace version.

Run `mise run install`, `mise run test`, and `mise run build`. The workspace's `go.work` resolves local modules. External consumers run `go get example.com/premise-module@v1.2.3`. Verify public modules with `GOWORK=off go test ./...` before release. Do not publish filesystem replace directives.

`ADDR` selects the bind address (default `:8080`). The server accepts Connect, gRPC, and gRPC-Web, including cleartext HTTP/2 for local gRPC clients. SIGINT starts graceful shutdown. Run `mise run generate` to regenerate checked-in code from proto/ with pinned plugins; no hosted generator is needed. Tests exercise all three protocols through a real HTTP/2 server.

The protobuf go_package option is deliberately stable. buf.gen.yaml maps the generated Go import path through both local plugins; this keeps registry module-path substitution out of length-encoded protobuf descriptors. Validation uses a different canonical module path to catch descriptor corruption.
