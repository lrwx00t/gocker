# gocker
DSL-like Dockerfile using Go


## Usage
```bash
❯ go run main.go
FROM ubuntu
CMD This is a sample comment
# This is a sample comment
RUN echo 'hello world'
MAINTAINER John Doe
EXPOSE 8080
EXPOSE 8080
RUN echo 'hello world'2
```