#!/bin/sh -e

VERSION=${1}
BINDIR=${2}

cd ${BINDIR}

for DIR in *
do
	DIR="$(basename ${DIR})"

	if echo "${DIR}" | grep -q "windows"
	then
		zip -9 -y -r -q "assert-v${VERSION}-${DIR}.zip" "${DIR}"
	else
		tar -czf "assert-v${VERSION}-${DIR}.tar.gz" "${DIR}"
	fi
done
