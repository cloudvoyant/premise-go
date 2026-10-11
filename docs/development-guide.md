# Development

Every directory under templates is a standalone Go module. Registry validation copies it to an isolated temporary directory, substitutes public and private-local policies, and runs every task with GOWORK=off and PREMISE_TEMPLATE_TEST=1. Run `mise run test` at the registry root.

Connect code is checked in. Run `mise run generate` in that template to regenerate with pinned local protobuf and Connect plugins. Repeat generation and compare gen/ to verify reproducibility.

The native Pulumi template uses Pulumi mock tests and a Go consumer under examples/consumer. It does not need Pulumi CLI or cloud providers. The optional provider template additionally uses the Pulumi CLI to generate SDKs and demonstrates Node consumption with mocks. Its SDK packaging is independent of native Go module release tags.

Pinned sources: [Go modules](https://go.dev/ref/mod), [Cobra](https://github.com/spf13/cobra/releases/tag/v1.10.2), [Connect Go](https://connectrpc.com/docs/go/getting-started/), [Pulumi Go provider SDK](https://www.pulumi.com/docs/iac/guides/building-extending/providers/sdks/), and [SDK generation](https://www.pulumi.com/docs/iac/cli/commands/pulumi_package_gen-sdk/). The optional provider uses pulumi-go-provider 1.2.0 and Pulumi CLI 3.200.0 as a compatible pair. Go 1.25.14, protobuf 1.36.12, Connect 1.21.0, htmx 2.0.8, and Cobra 1.10.2 are pinned.

The root test task also composes all six renamed modules in one temporary go.work and tests/vets them under shared dependency selection. Keep Pulumi pkg/v3 and sdk/v3 aligned at 3.200.0 across native/provider templates; mismatched versions can compile in isolation but fail together.
