package main

import (
	. "lrwx00t/gocker/generator"
)

// TODO:
// Formatter for single and double quotes
// Multiline
// runner/tester
// remaining docker keywords

func main() {
	test := `set -x \
	&& apk add --no-cache --virtual .build-deps \
		build-base \
		libffi-dev \
		openssl-dev \
	&& pip install --upgrade \
		--pre azure-cli \
		--extra-index-url https://azurecliprod.blob.core.windows.net/edge \
		--no-cache-dir \
	&& apk del .build-deps`
	GenerateDockerfile([]Dockerfile{
		FROM("ubuntu"),
		CMD("This is a sample comment"),
		COMMENT("This is a sample comment"),
		RUN([]string{"echo 'hello world'"}),
		MAINTAINER("John Doe"),
		ENTRYPOINT([]string{"ansible", "python"}),
		EXPOSE("8080"),
		EXPOSE("8081"),
		RUN([]string{"echo 'hello world'2"}),
		RUN([]string{test}),
		WORKDIR("/app"),
	})
}
