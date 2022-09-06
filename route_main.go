package main

import (
    "net/http"
    "fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello zengsiyu~")
}
