package main

import (
    "fmt"
    "github.com/gogf/greuse"
    "net/http"
    "os"
)

// We can create two processes with this code.
// Do some requests, then see the response from the server.
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
