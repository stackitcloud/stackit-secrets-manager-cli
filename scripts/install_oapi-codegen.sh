#!/usr/bin/env bash
set -e

VERSION=1.13.0
echo "installing oapi-codegen v${VERSION}"

TARGET_DIR=$(pwd)/bin
mkdir -p ${TARGET_DIR}
GOBIN=${TARGET_DIR} go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v${VERSION}
