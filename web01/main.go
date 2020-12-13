package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main(){
	addr := "0.0.0.0:8888"
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//请求行
		fmt.Println("method:", r.Method)
		fmt.Println("url:", r.URL)
		fmt.Println("protocol:", r.Proto)

		// 请求头
		// fmt.Println(r.Header)  // 请求头 内容
		header := r.Header
		fmt.Println(header.Get("User-Agent"))
		fmt.Println(header.Get("Token"))
		fmt.Println(r.Host) // 获取请求的地址
		fmt.Fprint(w, time.Now().Format("2020-09-09 12:09:02"))

		// 请求体
		fmt.Println("Body:")
		io.Copy(os.Stdout, r.Body)
	})
	http.ListenAndServe(addr, nil)
}



