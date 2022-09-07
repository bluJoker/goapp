package views

import (
    "goapp/models"
    "net/http"
    "encoding/json"
    "fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
    generateHTML(w, "", "web")
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
    first_user, err := models.GetFirstUser()
    if err != nil {
        fmt.Println(err)
    }
    output, err := json.MarshalIndent(&first_user, "", "\t\t")
    if err != nil {
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(output)

    //fmt.Fprintln(w, "Hello GetInfo...")
    return
}
