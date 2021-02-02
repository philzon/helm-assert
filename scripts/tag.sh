#!/bin/sh -e

COMMIT=$1

if [ -z "${COMMIT}" ]
then
	COMMIT="$(git rev-parse --short HEAD)"
fi

VERSION="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"

git tag --annotate "v${VERSION}" --message "Release ${VERSION}" "${COMMIT}"
