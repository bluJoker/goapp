package views

import (
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    generateHTML(w, "", "web")
}

