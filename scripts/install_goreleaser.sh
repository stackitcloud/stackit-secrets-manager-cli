#!/usr/bin/env bash
set -e

VERSION=1.18.2
echo "installing goreleaser v${VERSION}"

TARGET_DIR=$(pwd)/bin
mkdir -p ${TARGET_DIR}
GOBIN=${TARGET_DIR} go install github.com/goreleaser/goreleaser@v${VERSION}
