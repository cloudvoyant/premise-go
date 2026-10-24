# premise-htmx-app

Generated projects are packages in the workspace's root Go module. Keep the shared workspace module path when generating more projects. Applications stay under apps/ and libraries under libs/, each with its own Mise tasks. All projects use the same repository release version and one root Git tag.

The package.publication field in the workspace premise.yaml selects project publication tasks. private-local packages are still included in the shared module distribution. Use Go internal/ directories to restrict external imports.

Run `mise run install`, `mise run test`, and `mise run build`. The workspace's root `go.mod` resolves dependencies and local package imports. Install a released application with `go install example.com/premise-module/apps/premise-htmx-app@v1.2.3`. Verify public packages with `GOWORK=off go test ./...` before release. Do not publish filesystem replace directives.

`ADDR` selects the HTTP bind address (default `:8080`). Templates and htmx 2.0.8 are embedded in the binary, so deployment needs only the executable. SIGINT starts graceful shutdown. `POST /greet` returns an HTML fragment with escaped input. The vendored htmx script is distributed under its upstream 0BSD license.
