#!/bin/sh

set -u

# clean
FILES=$(find . -maxdepth 1 -type f ! -name "*.md" ! -name ".*" ! -name "*.sh" ! -name "config.yaml" ! -name "go.mod" ! -name "go.sum")
echo $FILES | xargs rm

# build
npx @openapitools/openapi-generator-cli generate -i https://raw.githubusercontent.com/traPtitech/traQ/master/docs/v3-api.yaml -g go -c config.yaml
go fmt ./...

rm -rf ./docs
