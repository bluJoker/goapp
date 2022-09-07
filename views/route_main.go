package views

import (
    "goapp/models"
    "net/http"
    "encoding/json"
    "fmt"
)

type userModel struct {
    Result    string `json:"result"`
    Name      string `json:"username"`
    Photo     string `json:"photo"`
}

func Index(w http.ResponseWriter, r *http.Request) {
    generateHTML(w, "", "web")
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
    first_user, err := models.GetFirstUser()
    if err != nil {
        fmt.Println(err)
    }
    user := userModel{}
    user.Name = first_user.Name
    user.Photo = first_user.Photo
    user.Result = "success"
    output, err := json.MarshalIndent(&user, "", "\t")
    if err != nil {
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(output)

    //fmt.Fprintln(w, "Hello GetInfo...")
    return
}
