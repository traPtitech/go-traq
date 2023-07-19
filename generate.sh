#!/bin/sh

set -eux

OPENAPI_GENERATOR_VERSION=$(cat ./.openapi-generator/VERSION)

# clean
mv README.md README.md.bak
cat ./.openapi-generator/FILES | xargs --no-run-if-empty rm -f

# fetch openapi-generator
NEEDS_FETCH='1'
if [ -e 'openapi-generator-cli.jar' ]; then
  VERSION_DOWNLOADED=`java -jar openapi-generator-cli.jar version`
  NEEDS_FETCH=`test $OPENAPI_GENERATOR_VERSION != $VERSION_DOWNLOADED && echo '1' || echo '0'`
fi
if [ $NEEDS_FETCH = '1' ]; then
  wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$OPENAPI_GENERATOR_VERSION/openapi-generator-cli-$OPENAPI_GENERATOR_VERSION.jar -O openapi-generator-cli.jar -q
fi

# build
java -jar openapi-generator-cli.jar generate \
  -i https://raw.githubusercontent.com/traPtitech/traQ/master/docs/v3-api.yaml \
  -g go \
  -c config.yaml
mv README.md client.md
mv README.md.bak README.md

# setup go
go fmt ./...
go mod tidy
