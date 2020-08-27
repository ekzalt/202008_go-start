# Go

[Documentation](https://golang.org/doc/)

[A Tour of Go](https://tour.golang.org/list)

## CLI commands

```bash
# help
go
go help

# version
go version

# run file
go run hello.go

# build file
go build hello.go

# cli doc
go doc fmt Printf

# run tests
go test
go test -v
```

## ENV

- GOPATH - project path
- GOROOT - Go path

## Project structure

- `bin/` - compiled files
- `pkg/` - service files
- `src/` - project code

## Management dependencies

```bash
# enable go mod
export GO111MODULE=on

# init go mod
go mod init
```

## Web Frameworks

- [Gin](https://github.com/gin-gonic/gin)
- [Mux](https://github.com/gorilla/mux)
