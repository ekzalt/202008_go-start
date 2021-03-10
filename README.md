# Go

- [Documentation](https://golang.org/doc/)
- [Packages](https://godoc.org/)
- [A Tour of Go](https://tour.golang.org/list)
- [Go by Example](https://gobyexample.com/)

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
# build with profiling
go build -gcflags=-m hello.go

# cli doc
go doc fmt Printf

# run tests
go test
go test -v
# run test by strict name
go test -run TestSumAndMult
# run tests by RegExp
go test -run TestSum -v

# go generate [util-name] [options]
# mark file
# //go:generate stringer -type=Pill
# execute generate
go generate path/to/original.go
# run generated
go run path/to/original.go path/to/generated.go
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

- [dep](https://github.com/golang/dep)
- [vgo](https://github.com/golang/go/wiki/vgo)

## Debug

[Delve](https://github.com/go-delve/delve) is a debugger for the Go programming language.

## Profiling

```go
// online diagnosis
import _ "net/http/pprof"
```

```bash
# profile code
go tool pprof -seconds 5 http://server/debuf/pprof/profile

# profile benchmark tests
go test -bench=.
# run benchmark with memory info
go test -bench=. -benchmem
# run benchmark with memory info and file output
go test -bench=YourBenchmarkNameHere -benchmem -memprofile=mem.out
# run benchmark with cpu info output
go test -bench=YourBenchmarkNameHere -cpuprofie cpu.out
go test -bench=. -cpuprofie cpu.out

go tool pprof -alloc_space your_compiled_binary_file mem.out
go tool pprof your_compiled_binary_file cpu.out

list
list yourFuncName
```

## Web Frameworks

- [FastHTTP](https://github.com/valyala/fasthttp) is the fastest http lib, [documentation](https://godoc.org/github.com/valyala/fasthttp)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorilla web toolkit](https://www.gorillatoolkit.org/) and [Mux](https://github.com/gorilla/mux)
- [Beego](https://beego.me/)

## ORM

- [GORM](https://github.com/go-gorm/gorm), [documentation](https://gorm.io/)

## WebSocket

- [WebSocket](https://github.com/gorilla/websocket)

## JSON

- [EasyJson](https://github.com/mailru/easyjson)

```bash
# install package
go get -u github.com/mailru/easyjson/...
# run generate Marshal/Unmarshal methods for structs
easyjson -all ./path/to/file.go
```

## gRPC

Download Protobuf Compiler [protoc](https://github.com/protocolbuffers/protobuf/releases)
