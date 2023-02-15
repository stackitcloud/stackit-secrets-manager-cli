#!/usr/bin/env bash
set -e

VERSION=1.51.1
echo "installing golangci-lint v${VERSION}"

echo "operating system: ${OSTYPE}"
TYPE=windows
if [[ "${OSTYPE}" == linux* ]]; then
  TYPE=linux
elif [[ "${OSTYPE}" == darwin* ]]; then
  TYPE=darwin
fi

echo "cpu architecture: $(uname -m)"
case $(uname -m) in
arm64)
    ARCH=arm64
    ;;
*)
    ARCH=amd64
    ;;
esac

mkdir -p bin/
curl \
  --output bin/golangci-lint.tar.gz \
  --location \
  --silent \
  https://github.com/golangci/golangci-lint/releases/download/v${VERSION}/golangci-lint-${VERSION}-${TYPE}-${ARCH}.tar.gz

if [[ "${OSTYPE}" == linux* ]]; then
  tar \
    -x \
    --directory bin \
    --gunzip \
    --strip-components 1 \
    --file bin/golangci-lint.tar.gz \
    --wildcards **/golangci-lint
elif [[ "${OSTYPE}" == darwin* ]]; then
  tar \
    -x \
    --directory bin \
    --gunzip \
    --strip-components 1 \
    --file bin/golangci-lint.tar.gz \
    --include '*/golangci-lint'
fi

rm bin/golangci-lint.tar.gz
