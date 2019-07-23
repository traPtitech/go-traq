#!/bin/sh

set -u

# clean
FILES=$(find . -maxdepth 1 -type f ! -name "*.md" ! -name ".*" ! -name "*.sh" ! -name "go.mod" ! -name "go.sum")
echo $FILES | xargs rm

# build
npx @openapitools/openapi-generator-cli generate -i https://raw.githubusercontent.com/traPtitech/traQ/master/docs/swagger.yaml -g go
go fmt ./...

rm -rf ./docs
