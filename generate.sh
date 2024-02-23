#!/bin/sh

set -eux

OPENAPI_GENERATOR_VERSION=$(cat ./.openapi-generator/VERSION)

# clean
mv README.md README.md.bak
cat ./.openapi-generator/FILES | xargs -r rm -f

# build
docker run --rm -v "${PWD}:/local" -u $(id -u) openapitools/openapi-generator-cli:v$OPENAPI_GENERATOR_VERSION generate \
  -i https://raw.githubusercontent.com/traPtitech/traQ/master/docs/v3-api.yaml \
  -g go \
  -c /local/config.yaml \
  -o /local
mv README.md client.md
mv README.md.bak README.md

# setup go
go fmt ./...
go mod tidy
