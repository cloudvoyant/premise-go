# premise-go

Go templates for Premise: Cobra CLI, Connect/gRPC service, HTMX app, Go libraries, native Go Pulumi components, and an optional cross-language Pulumi provider.

Generated projects keep the standard apps/<name> and libs/<name> layout with per-project Mise tasks. They share one root go.mod, go.sum, dependency set, and repository release tag. There is no go.work or per-library module tag. External clients can import an individual library package at the shared module version.

Generate with `pm generate cloudvoyant/premise-go:premise-go-lib`, or another selector in premise.yaml. Supply the same workspace Go module path for every project. During tandem development, use an explicit local registry source with the matching Premise branch.

Run `mise install`, then `mise run test` for the same `premise template test` contract used by the other registries. Run `mise run e2e` for real generation of all six templates and shared-module tests. Set `PREMISE_BIN=/absolute/path/to/premise` to test a local core build. Mise selects the exact pinned CLI by default. The optional Pulumi provider also runs its Node consumer through the ordinary `test` task.

The first generated Go project copies `.golangci.yml` and `.editorconfig` to the workspace root when those files are missing. Later projects keep the client rules. All templates use those shared rules for formatting, imports, static analysis, error handling, and a cyclomatic complexity limit of 30. The standard Premise commit and merge workflows run the registry tasks and template contracts.

The `package.public` boolean controls project publication tasks and defaults to true. It does not exclude source from the shared root module. The optional template that generates bindings for other languages can be removed independently.
