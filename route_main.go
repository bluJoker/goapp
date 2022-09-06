package main

import (
    "net/http"
    "fmt"
)

func test(w http.ResponseWriter, r *http.Request) {
    line1 := `<h1 style="text-align: center">Hello zengsiyu~</h1>`
    line2 := `<p style="text-align: center;">
    <img src="https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fpic2.zhimg.com%2F80%2Fv2-7a68d835151a360a121fb6490f53367e_720w.jpg%3Fsource%3D1940ef5c&refer=http%3A%2F%2Fpic2.zhimg.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1665021505&t=d023f7b25e0fd4f994803c675d65005c">
    </p>`
    line3 := "<hr>"
    fmt.Fprintln(w, line1, line3, line2)
}
