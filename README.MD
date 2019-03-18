# greuseport [![Go Doc](https://godoc.org/github.com/gogf/greuseport?status.svg)](https://godoc.org/github.com/gogf/greuseport) [![Documents](https://img.shields.io/badge/docs-100%25-green.svg)](https://goframe.org) [![License](https://img.shields.io/github/license/gogf/greuseport.svg?style=flat)](https://github.com/gogf/greuseport) [![Language](https://img.shields.io/badge/language-go-blue.svg)](https://github.com/gogf/greuseport) [![Release](https://img.shields.io/github/release/gogf/greuseport.svg?style=flat)](https://github.com/gogf/greuseport/releases)

Package greuseport provides Listen and Dial functions that set socket options in order to be able to reuse ports. You should only use this package if you know what SO_REUSEADDR and SO_REUSEPORT are.

# Installation
```
go get -u github.com/gogf/greuseport
```
or use `go.mod`
```
require github.com/gogf/greuseport latest
```

# Documentation

* [https://godoc.org/github.com/gogf/greuseport](https://godoc.org/github.com/gogf/greuseport)

# Quick Start
```go
package main

import (
    "fmt"
    "github.com/gogf/greuseport"
    "net/http"
    "os"
)

// We can create two processes with this code.
// Do some requests, then watch the output of the console.
func main() {
    listener, err := greuseport.Listen("tcp", ":8881")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    server := &http.Server{}
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "gid: %d, pid: %d\n", os.Getgid(), os.Getpid())
    })

    panic(server.Serve(listener))
}
```


# License

`greuseport` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.

