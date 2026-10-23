#!/bin/sh
set -eu
version="${RELEASE_VERSION:-0.1.0}"
version="${version#v}"
mkdir -p dist
# SDK generation replaces its output directory. Preserve dependency selections.
rm -f dist/node-sdk-package-lock.json
if [ -f sdk/nodejs/package-lock.json ]; then
    cp sdk/nodejs/package-lock.json dist/node-sdk-package-lock.json
fi
go build -ldflags="-X main.version=$version" -o dist/pulumi-resource-premise-pulumi-component-package .
pulumi package get-schema ./dist/pulumi-resource-premise-pulumi-component-package > schema.json
pulumi package gen-sdk schema.json --language go --out sdk
pulumi package gen-sdk schema.json --language nodejs --out sdk
python3 scripts/pin-sdk.py
if [ -f dist/node-sdk-package-lock.json ]; then
    cp dist/node-sdk-package-lock.json sdk/nodejs/package-lock.json
fi
(cd sdk/nodejs && npm install --ignore-scripts --no-audit --no-fund)
