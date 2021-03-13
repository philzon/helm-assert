#!/bin/sh -e

NAME="$(cat plugin.yaml | grep "name" | cut -d '"' -f 2)"
VERSION="$(curl --silent https://api.github.com/repos/philzon/helm-assert/releases/latest | grep tag_name | sed -E 's/.*"([^"]+)".*/\1/')"
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

URL="https://github.com/philzon/helm-assert/releases/download/${VERSION}/${NAME}-${VERSION}-${KERNEL}-${ARCH}.tar.gz"

echo "${URL}"

mkdir --parent "bin"
mkdir --parent "releases/${VERSION}"

if [ -x "$(which curl 2>/dev/null)" ]
then
	curl -sSL "${URL}" -o "releases/${VERSION}.tar.gz"
else
	wget -q "${URL}" -O "releases/${VERSION}.tar.gz"
fi

tar -xzf "releases/${VERSION}.tar.gz" -C "releases/${VERSION}"
mv "releases/${VERSION}/${NAME}" "bin/${NAME}" || \
	mv "releases/${VERSION}/${NAME}.exe" "bin/${NAME}"
mv "releases/${VERSION}/plugin.yaml" .
