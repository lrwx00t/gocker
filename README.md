[![Go Reference](https://pkg.go.dev/badge/github.com/lrwx00t/gocker.svg)](https://pkg.go.dev/github.com/lrwx00t/gocker)

# gocker
DSL-like Dockerfile using Go

[Go Playground](https://go.dev/play/p/x7ShKTcTrx3)

## Usage

After installing the package:
```
go get -u github.com/lrwx00t/gocker/generator
```

Simple `main.go`:
```go
package main

import (
	. "github.com/lrwx00t/gocker/generator"
)

func main() {
	GenerateDockerfile([]Dockerfile{
		FROM("ubuntu"),
		CMD("This is a sample comment"),
		EXPOSE("8080"),
		EXPOSE("8081"),
		WORKDIR("/app"),
	})
}
```

Output:

```bash
❯ go run main.go
FROM ubuntu
CMD This is a sample comment
EXPOSE 8080
EXPOSE 8081
WORKDIR /app
```