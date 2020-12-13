package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	addr := ":8080"
	// 动态的生成响应结果
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// html
		f, _ := os.Open("./template/index.html")
		io.Copy(writer, f)
	})
	http.ListenAndServe(addr, nil)
}
