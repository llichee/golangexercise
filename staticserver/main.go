package main

import "net/http"

func main() {
	addr := ":8989"
	// http.Dir 类型转换的功能 ==> type Dir string
//	http.Handle("/static/", http.FileServer(http.Dir(".")))   在static目录下找数据 然后返给浏览器
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("./static/"))))   //访问 www 目录跳转至static目录下


	http.ListenAndServe(addr, nil)
}
