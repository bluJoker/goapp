package views

import (
    "goapp/models"
    "net/http"
    "fmt"
    "encoding/json"
)

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

// GET /logout
// Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
    cookie, err := request.Cookie("_cookie")
    if err != http.ErrNoCookie {
        //warning(err, "Failed to get cookie")
        session := models.Session{Uuid: cookie.Value}
        session.DeleteByUUID()
    }

    result := resultCode{}
    result.Result = "success"
    output, err := json.MarshalIndent(&result, "", "\t")
    if err != nil {
        return
    }
    writer.Header().Set("Content-Type", "application/json")
    writer.Write(output)
}

// POST /signup
// Create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm()
    if err != nil {
        //danger(err, "Cannot parse form")
    }
    username := request.PostFormValue("username")
    email := request.PostFormValue("email")
    password := request.PostFormValue("password")
    password_confirm := request.PostFormValue("password_confirm")
    fmt.Println("signup: ", username, email, password, password_confirm)

    if username == "" || password == "" {
        JsonResponse(writer, "用户名和密码不能为空")
        return
    }
    if password != password_confirm {
        JsonResponse(writer, "两次密码不一致")
        return
    }
    _, err = models.UserByEmail(request.PostFormValue("email"))
    if err == nil {
        JsonResponse(writer, "该邮箱已注册")
        return
    }
    new_user := models.User{
        Name:     request.PostFormValue("username"),
        Email:    request.PostFormValue("email"),
        Password: request.PostFormValue("password"),
        Photo:    request.PostFormValue("photo"),
    }
    if err := new_user.Create(); err != nil {
        //danger(err, "Cannot create user")
    }
    //http.Redirect(writer, request, "/login", 302) // 注册成功后转到登录界面
    JsonResponse(writer, "success")
    return
}
