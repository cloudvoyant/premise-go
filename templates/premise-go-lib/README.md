# premise-go-lib

Generated projects are packages in the workspace's root Go module. Keep the shared workspace module path when generating more projects. Applications stay under apps/ and libraries under libs/, each with its own Mise tasks. All projects use the same repository release version and one root Git tag.

The package.publication field in the workspace premise.yaml selects project publication tasks. private-local packages are still included in the shared module distribution. Use Go internal/ directories to restrict external imports.

Run `mise run install`, `mise run test`, and `mise run build`. The workspace's root `go.mod` resolves dependencies and local package imports. External consumers run `go get example.com/premise-module/libs/premise-go-lib@v1.2.3`. Verify public packages with `GOWORK=off go test ./...` before release. Do not publish filesystem replace directives.
