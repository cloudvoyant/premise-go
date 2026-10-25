# premise-pulumi-component-package

An optional Pulumi component provider with an executable and generated Go and Node SDKs. Use `premise-pulumi-component-lib` when consumers only need native Go components.

## Use

From a separate Go Pulumi project, fetch the SDK package:

```bash
go get example.com/premise-module/libs/premise-pulumi-component-package/sdk/go/greeting@v1.2.3
```

Import `example.com/premise-module/libs/premise-pulumi-component-package/sdk/go/greeting`. The SDK belongs to the shared workspace module and uses its repository tag.

## Develop

```bash
mise run install
mise run format:check
mise run lint
mise run test
mise run build
```

Formatting and lint tasks use the workspace's root `.golangci.yml`.

## Generate SDKs

Run `mise run generate` to build the provider, extract its schema, and regenerate the pinned Go and Node SDKs. Run `mise run test:sdk` for the Node consumer with Pulumi mocks. The normal test task also includes that consumer. Go tests cover the Go SDK consumer. Neither consumer deploys cloud resources.

Before a shared release, regenerate the provider and SDKs at the shared version:

```bash
RELEASE_VERSION=1.2.3 mise run generate
```

A leading `v` is accepted. Commit the regenerated source and lockfile before release. Generation without `RELEASE_VERSION` restores development artifacts at `0.1.0`.

## Distribution

The provider executable uses development version `0.1.0`. A release build replaces it through `-X main.version=<shared-version>`. Its resource token is `premise-pulumi-component-package:index:Greeting`.

Before distributing binary plugins or Node packages, configure the namespace, plugin download URL, SDK package name, package feeds, and release tasks. Node and plugin distribution remain separate from the Go module tag. `package.main` and `package.binary` in `premise.yaml` let the coordinated release include the provider executable.

## Workspace

This project is a package in the workspace's root Go module. Keep the same module path when you add projects. Applications live under `apps/`, and libraries live under `libs/`. All packages share dependencies and one repository tag.

`package.public: false` skips project publication tasks. It does not remove source from the module. Use Go `internal/` directories to restrict external imports. Before release, run `GOWORK=off go test ./...` from the workspace root and remove filesystem `replace` directives.
