package views

import (
    "net/http"
    "html/template"
    "fmt"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
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
