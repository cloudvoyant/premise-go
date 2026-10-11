# premise-go

Go templates for premise: Cobra CLI, Connect/gRPC service, HTMX app, separately consumable Go libraries, native Go Pulumi components, and an optional cross-language Pulumi component provider.

Generate with `pm generate cloudvoyant/premise-go:premise-go-lib` (or another selector in premise.yaml). During tandem development use an explicit local registry source. Projects use separate modules and one root go.work; public modules share the workspace release version through prefixed Git tags. No cloud deployment is required.

Run `mise run test` to validate template modules. The cross-language template adds provider packaging and SDK generation and can be removed independently of native components.
