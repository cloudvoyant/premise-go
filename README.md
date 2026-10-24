# premise-go

Go templates for Premise: Cobra CLI, Connect/gRPC service, HTMX app, Go libraries, native Go Pulumi components, and an optional cross-language Pulumi provider.

Generated projects keep the standard apps/<name> and libs/<name> layout with per-project Mise tasks. They share one root go.mod, go.sum, dependency set, and repository release tag. There is no go.work or per-library module tag. External clients can import an individual library package at the shared module version.

Generate with `pm generate cloudvoyant/premise-go:premise-go-lib`, or another selector in premise.yaml. Supply the same workspace Go module path for every project. During tandem development, use an explicit local registry source with the matching Premise branch.

Run `mise install` and `mise run test` to validate every standalone template contract, then generate all six projects through the pinned Premise CLI and test their shared root module. To test a local core checkout, set `PREMISE_BIN=/absolute/path/to/premise` when running `mise run test`. The workspace check uses real generation; it does not construct a parallel module merge. The package.public boolean controls project publication tasks and defaults to true. It does not exclude source from the root module. The optional cross-language template can be removed independently.
