package main

import (
	. "github.com/lrwx00t/gocker/generator"
)

// TODO:
// Formatter for single and double quotes
// Multiline
// runner/tester
// remaining docker keywords

func Example() {
	GenerateDockerfile([]Dockerfile{
		FROM("ubuntu"),
		CMD("This is a sample comment"),
		EXPOSE("8080"),
		EXPOSE("8081"),
		WORKDIR("/app"),
	})
	// Output:
	// FROM ubuntu
	// CMD This is a sample comment
	// EXPOSE 8080
	// EXPOSE 8081
	// WORKDIR /app
}
