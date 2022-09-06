package main

import (
    "fmt"
    "log"
    "encoding/json"
    "os"
    "time"
    "net/http"
    "html/template"
)

type Configuration struct {
    Address string
    Static  string
}

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
    fmt.Println(a)
}

func init() {
    loadConfig()
    file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file", err)
    }
    logger = log.New(file, "INFOW: ", log.Ldate|log.Ltime|log.Lshortfile)
    //warning(err, "Failed to get cookie")
}

func loadConfig() {
    file, err := os.Open("config.json")
    if err != nil {
        log.Fatalln("Cannot open config file", err)
    }
    decoder := json.NewDecoder(file)
    config = Configuration{}
    err = decoder.Decode(&config)
    if err != nil {
        log.Fatalln("Cannot get configuration from file", err)
    }

}


//拦截器返回一个函数供调用，在这个函数里添加自己的逻辑判断即可 h(w,r)及是调用用户自己的处理函数。h 是函数指针
func handleInterceptor(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("[%s] \"%s %s %s\" %d [%s]\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.Proto, r.ContentLength, r.RemoteAddr)
        h(w, r)
    }
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
    var files []string
    for _, file := range filenames {
        files = append(files, fmt.Sprintf("templates/multiends/%s.html", file))
    }

    // 可变参数变量是一个包含所有参数的切片，如果要将这个含有可变参数的变量传递给下一个可变参数函数，可以在传递时给可变参数变量后面添加...，
    // 这样就可以将切片中的元素进行传递，而不是传递可变参数变量本身。
    templates := template.Must(template.ParseFiles(files...))

    // 执行模板，将data数据和layout模板融合生成最终的HTML内容
    // layout模板不等于layout.html文件，而是通过{{ define "layout" }}和{{ end }}之间包含的内容
    templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
func info(args ...interface{}) {
    logger.SetPrefix("INFO ")
    logger.Println(args...)
}

func danger(args ...interface{}) {
    logger.SetPrefix("ERROR ")
    logger.Println(args...)
}

func warning(args ...interface{}) {
    logger.SetPrefix("WARNING ")
    logger.Println(args...)
}

// version
func version() string {
    return "0.1"
}

