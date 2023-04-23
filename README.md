# gocker
DSL-like Dockerfile using Go


## Usage

After installing the package:
```
go get -u lrwx00t/gocker/generator
```

Simple `main.go`:
```go
package main

import (
	. "lrwx00t/gocker/generator"
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