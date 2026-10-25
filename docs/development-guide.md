# Development

Source templates have standalone dependency manifests so their contracts can run in isolation. Their module paths include the intended apps/ or libs/ project path. During generation, the Go provider merges dependencies into the root workspace manifest and removes the project manifests. Native Go install tasks regenerate checksums. Generated projects therefore contain packages, with no nested go.mod and no go.work.

The root mise.toml declares a monorepo and discovers tasks under templates/premise-*, apps/*, and libs/*. Run `mise tasks --all` to list tasks. Each template owns its tools and tasks. The registry root mise.toml is not copied into generated clients.

Run `mise run test` at the registry root for the common Premise template contract. Run `mise run e2e` for the Go integration suite. It generates all six projects into a shared module, preserves client quality rules, and rejects formatting and complexity violations. Run one source template directly with `mise run //templates/premise-cobra-cli:test`.

Connect code is checked in. Run `mise run generate` in that template to regenerate with pinned local protobuf and Connect plugins. Repeat generation and compare gen/ to verify reproducibility.

The native Pulumi template uses Pulumi mock tests and a Go consumer under examples/consumer. It does not need Pulumi CLI or cloud providers. The optional provider template additionally uses the Pulumi CLI to generate SDKs and demonstrates Node consumption with mocks. Its SDK packaging is independent of native Go module release tags.

Pinned sources: [Go modules](https://go.dev/ref/mod), [Cobra](https://github.com/spf13/cobra/releases/tag/v1.10.2), [Connect Go](https://connectrpc.com/docs/go/getting-started/), [Pulumi Go provider SDK](https://www.pulumi.com/docs/iac/guides/building-extending/providers/sdks/), and [SDK generation](https://www.pulumi.com/docs/iac/cli/commands/pulumi_package_gen-sdk/). The optional provider uses pulumi-go-provider 1.2.0 and Pulumi CLI 3.200.0 as a compatible pair. Go 1.25.14, protobuf 1.36.12, Connect 1.21.0, htmx 2.0.8, and Cobra 1.10.2 are pinned.

The shared-module test checks one root go.mod and no go.work. Keep Pulumi pkg/v3 and sdk/v3 versions aligned across the native and provider templates. All projects share dependency versions in a generated workspace.

Template package settings live in premise.yaml beside each template declaration. Premise applies questionnaire substitutions and records the settings in the generated workspace project. There are no Go sidecar files. Core generation tests cover true and false publication settings.

The registry uses the standard Premise commit and merge workflows. Root Mise tasks own registry checks, and each template supplies the same contract tasks as Cargo and Bun. The separate Python test runner is removed.

Workspace quality rules live in `.golangci.yml` and `.editorconfig`. The registry declares them in `workspace_init_files`, so generation copies missing files without replacing existing client configuration. Linters include error handling, static analysis, HTTP response cleanup, and cyclomatic complexity. Formatting uses gofmt and goimports. Clients can edit their root rules after the first project.
