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

# improve time format from second to nanosec for traQgazer
# https://github.com/traP-jp/traQ-gazer/issues/138
# changed at https://github.com/traPtitech/go-traq/pull/8
sed -i.bak s/\(time.RFC3339\)/\(time.RFC3339Nano\)/g ./client.go
rm ./client.go.bak

# openapi-generator bug: https://github.com/OpenAPITools/openapi-generator/issues/20749
# Replace default value of time.Time from "0000-01-01T00:00Z" (string) to time.Time (zero value)
find . -name "*.go" -print0 | xargs -0 sed -i -E 's/var\s+defaultValue\s+time\.Time\s*=\s*"0000-01-01T00:00Z"/var defaultValue time.Time/g'

# setup go
go fmt ./...
go mod tidy
