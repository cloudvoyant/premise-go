# premise-pulumi-component-package

Each project is its own Go module. Choose a canonical module path under your repository and a publication policy. `private-local` modules stay in the workspace; `public` modules receive directory-prefixed tags at the shared workspace version.

This optional template is a component provider with an executable and generated SDKs.
The root Go package builds the provider executable.
Go consumers import `example.com/premise-module/sdk/go/greeting`, the generated SDK package within the same module.

In the provider project, install dependencies, run tests, and build the executable:

```sh
mise run install
mise run test
mise run build
```

In a separate Go Pulumi project, fetch the published SDK package:

```sh
go get example.com/premise-module/sdk/go/greeting@v1.2.3
```

The workspace's `go.work` resolves local modules during development.
Verify public modules with `GOWORK=off go test ./...` before release.
Do not publish filesystem replace directives.

Run `mise run generate` to build the provider, extract its schema, and regenerate pinned Go and Node SDKs.
Run `mise run test:sdk` for a local Node SDK consumer with Pulumi mocks.
The Go tests also check the generated Go SDK consumer.
Neither demonstration deploys cloud resources.

The provider uses development version `0.1.0`, replaced through `-X main.version=<shared-version>` when building a release, and fixed token (`premise-pulumi-component-package:index:Greeting`). Before distributing binary plugins or Node packages, configure your own namespace, plugin download URL, SDK package name, package feeds, and release automation. The Go SDK can use the workspace's coordinated module tags; npm/plugin distribution remains an explicit separate packaging step. This exploratory template can be removed before merge without changing the native Go component library.

Generated Go SDK packages live inside this root module; there is no nested go.mod or SDK-specific tag requirement. The sidecar's main and binary fields let the coordinated release package the provider executable alongside the published module. Node package/plugin publication is intentionally a separate opt-in step.

Before freezing a shared release commit, run `RELEASE_VERSION=1.2.3 mise run generate` (a leading `v` is accepted). This builds the provider at the shared version and regenerates the schema and both SDKs with matching versions, including SDK plugin metadata. Commit the regenerated source and lockfile with the coordinated release preparation. `mise run generate` without RELEASE_VERSION restores development artifacts at `0.1.0`. No independent package version is calculated.
