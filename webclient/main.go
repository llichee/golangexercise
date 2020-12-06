package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 模拟浏览器发送请求
	response, err := http.Get("http://localhost:8888/help")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Proto)
	fmt.Println(response.StatusCode)
	io.Copy(os.Stdout, response.Body)

	req, err := http.NewRequest("DELETE", "http://localhost:8888/home", nil)
	// 创建一个客户端，用来发起请求
	client := http.Client{}
	// 用客户端发起请求
	response, err = client.Do(req)

	io.Copy(os.Stdout, response.Body)

	response_new, err := http.Get("http://www.baidu.com")
	io.Copy(os.Stdout, response_new.Body)
}
