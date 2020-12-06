package main

import (
	"fmt"
	"net/http"
)

// 定义一个 名为 Home 的处理器函数 ======
func Home(w http.ResponseWriter, r *http.Request) {
	// 用户提交的数据 http内容 -> go代码转换 http.Request
	w.Write([]byte("hello world"))
	fmt.Println("xxx")
}

// 自定义结构体 +++++
type Help struct {
}
// 自定义结构体方法 ++++
func (h *Help) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "help page")
}
func main() {

	addr := ":8888"
	// url => 找处理器函数 => 调用处理器函数(http包)
	// 指定url和处理器关系
	// 处理器函数签名由 http 包定义 =====
	http.HandleFunc("/home", Home)

	// 自定义结构体后 如下方式调用 ++++++
	http.Handle("/help", new(Help))

	// 启动服务
	http.ListenAndServe(addr, nil)
}
