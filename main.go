package main

import (
    "net/http"
)

func main() {
    p("App", version(), "started at", config.Address)

    mux := http.NewServeMux()
    mux.HandleFunc("/test", test)

    server := &http.Server{
        Addr: config.Address,
        Handler: mux,
        MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()
}
