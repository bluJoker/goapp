package main

import (
    "goapp/views"
    "net/http"
)

func main() {
    p("App", version(), "started at", config.Address)

    mux := http.NewServeMux()

    // 指定静态文件解析路径
    files := http.FileServer(http.Dir(config.Static))
    mux.Handle("/static/", http.StripPrefix("/static", files))

    mux.HandleFunc("/test", handleInterceptor(test))

    // urls
    // defined in views/route_main.go
    mux.HandleFunc("/", handleInterceptor(views.Index))
    mux.HandleFunc("/settings/getinfo", handleInterceptor(views.GetInfo))
    mux.HandleFunc("/settings/login", handleInterceptor(views.Authenticate))
    mux.HandleFunc("/settings/logout", handleInterceptor(views.Logout))
    mux.HandleFunc("/settings/signup", handleInterceptor(views.SignupAccount))

    server := &http.Server{
        Addr: config.Address,
        Handler: mux,
        MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()
}
