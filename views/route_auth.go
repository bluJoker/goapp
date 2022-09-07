package views

import (
    "goapp/models"
    "net/http"
    "fmt"
    "encoding/json"
)

type resultCode struct {
    Result string `json:"result"`
}

type userModel struct {
    Result    string `json:"result"`
    Email      string `json:"email"`
    Photo     string `json:"photo"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
    session, err := session(w, r)
    if err == nil {
        fmt.Println("sessionw2: ", session)

        user, err := models.UserByEmail(session.Email)
        login_data := userModel{}
        login_data.Email = user.Email
        login_data.Photo = user.Photo
        login_data.Result = "success"
        login_output, err := json.MarshalIndent(&login_data, "", "\t")
        if err != nil {
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(login_output)
    } else {
        logout_data := userModel{}
        logout_data.Result = "用户未登录"
        logout_output, err := json.MarshalIndent(&logout_data, "", "\t")
        if err != nil {
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(logout_output)
    }

    //fmt.Fprintln(w, "Hello GetInfo...")
    return
}

// POST /authenticate
// Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm() // 获取表单数据
    fmt.Println(request.PostForm)
    user, err := models.UserByEmail(request.PostFormValue("email"))
    if err != nil {
        //danger(err, "Cannot find user")
        fmt.Println("Cannot find user!~")
    }
    if user.Password == models.Encrypt(request.PostFormValue("password")) {
        session, err := user.CreateSession()
        if err != nil {
            //danger(err, "Cannot create session")
        }
        cookie := http.Cookie{
            Name:     "_cookie",
            Value:    session.Uuid,
            HttpOnly: true,
        }
        http.SetCookie(writer, &cookie)

        result := resultCode{}
        result.Result = "success"
        output, err := json.MarshalIndent(&result, "", "\t")
        if err != nil {
            return
        }
        writer.Header().Set("Content-Type", "application/json")
        writer.Write(output)
    } else {
        result := resultCode{}
        result.Result = "用户名或密码不正确"
        output, err := json.MarshalIndent(&result, "", "\t")
        if err != nil {
            return
        }
        writer.Header().Set("Content-Type", "application/json")
        writer.Write(output)
    }
}
