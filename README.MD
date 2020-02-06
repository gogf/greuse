# greuse

Package greuse provides Listen and Dial functions that set socket options in order to be able to reuse ports. 
You should only use this package if you know what SO_REUSEADDR and SO_REUSEPORT are.

# Installation
```
go get -u -v github.com/gogf/greuse
```
or use `go.mod`
```
require github.com/gogf/greuse latest
```

# Documentation

* [https://godoc.org/github.com/gogf/greuse](https://godoc.org/github.com/gogf/greuse)

# Quick Start
```go
package main

import (
    "fmt"
    "github.com/gogf/greuse"
    "net/http"
    "os"
)

// We can create two processes with this code.
// Do some requests, then watch the output of the console.
func main() {
    listener, err := greuse.Listen("tcp", ":8881")
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

`greuse` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.

