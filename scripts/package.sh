#!/bin/sh -e

VERSION="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
BINDIR=${1}

cd $BINDIR

for DIR in *
do
	BASE="$(basename ${DIR})"

	cd ${BASE}

	if echo "$PWD" | grep -q "windows"
	then
		zip -9 -y -r -q "assert-v${VERSION}-${BASE}.zip" *
		mv "assert-v${VERSION}-${BASE}.zip" ../
	else
		tar -czf "assert-v${VERSION}-${BASE}.tar.gz" *
		mv "assert-v${VERSION}-${BASE}.tar.gz" ../
	fi

	cd ..
done
