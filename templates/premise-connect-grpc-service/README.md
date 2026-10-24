# premise-connect-grpc-service

Generated projects are packages in the workspace's root Go module. Keep the shared workspace module path when generating more projects. Applications stay under apps/ and libraries under libs/, each with its own Mise tasks. All projects use the same repository release version and one root Git tag.

The package.public field in the workspace premise.yaml selects project publication tasks and defaults to true. Set it to false to skip those tasks. All packages remain in the shared module distribution. Use Go internal/ directories to restrict external imports.

Run `mise run install`, `mise run test`, and `mise run build`. The workspace's root `go.mod` resolves dependencies and local package imports. Install a released application with `go install example.com/premise-module/apps/premise-connect-grpc-service@v1.2.3`. Verify public packages with `GOWORK=off go test ./...` before release. Do not publish filesystem replace directives.

`ADDR` selects the bind address (default `:8080`). The server accepts Connect, gRPC, and gRPC-Web, including cleartext HTTP/2 for local gRPC clients. SIGINT starts graceful shutdown. Run `mise run generate` to regenerate checked-in code from proto/ with pinned plugins; no hosted generator is needed. Tests exercise all three protocols through a real HTTP/2 server.

The protobuf go_package option is deliberately stable. buf.gen.yaml maps the generated Go import path through both local plugins; this keeps registry module-path substitution out of length-encoded protobuf descriptors. Validation uses a different canonical module path to catch descriptor corruption.
