#!/bin/sh -e

NAME="$(cat plugin.yaml | grep "name" | cut -d '"' -f 2)"
VERSION="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
KERNEL="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

if [ "${KERNEL}" != "linux" ] && [ "${KERNEL}" != "darwin" ]
then
	KERNEL="windows"
fi

if [ "${ARCH}" != "aarch64" ] && [ "${ARCH}" != "arm64" ]
then
	ARCH="amd64"
fi

URL="https://github.com/philzon/helm-assert/releases/download/${VERSION}/assert-v${VERSION}-${KERNEL}-${ARCH}.tar.gz"

echo "${URL}"

mkdir --parent "bin"
mkdir --parent "releases/v${VERSION}"

if [ -x "$(which curl 2>/dev/null)" ]
then
	curl -sSL "${URL}" -o "releases/v${VERSION}.tar.gz"
else
	wget -q "${URL}" -O "releases/v${VERSION}.tar.gz"
fi

tar -xzf "releases/v${VERSION}.tar.gz" -C "releases/v${VERSION}"
mv "releases/v${VERSION}/${NAME}" "bin/${NAME}" || \
	mv "releases/v${VERSION}/${NAME}.exe" "bin/${NAME}"
mv "releases/v${VERSION}/plugin.yaml" .
