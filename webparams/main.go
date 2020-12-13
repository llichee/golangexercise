package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main(){
	addr := "0.0.0.0:8888"
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(strings.Repeat("+", 30))
		fmt.Println(r.Method)
		// 1. 提交数据方式
		/*
			在url中传递数据 ==> url?argname1=argvalue1&argname2=argvalue2
		 */
		//fmt.Println(r.URL)
		r.ParseForm() // 解析参数
		fmt.Println(r.Form)  // 接收的参数类型都是string
		fmt.Println(r.Form.Get("x"))  // 只返回第一个参数
		fmt.Println(r.Form["x"])  // 返回参数对应的字符串切片
		// 以上Form 可以获取URL中参数也可以获取Body中参数

		// 2. 通过body提交数据
		/*
			body 中数据格式：
			第一种格式(application/x-www-form-urlencoded)： a=b&c=d

			第三种格式(multipart/form-data)
			其他格式：
				第二种格式(application/json)
					(application/xml)
		 */

		fmt.Println(r.PostForm) // PostForm 只可获取Body中的参数


		fmt.Fprint(w, time.Now().Format("2020-09-09 12:09:02"))

		//请求行
		//fmt.Println("method:", r.Method)
		//fmt.Println("url:", r.URL)
		//fmt.Println("protocol:", r.Proto)

		// 请求头
		// fmt.Println(r.Header)  // 请求头 内容
		//header := r.Header
		//fmt.Println(header.Get("User-Agent"))
		//fmt.Println(header.Get("Token"))
		//fmt.Println(r.Host) // 获取请求的地址


		// 请求体
		//fmt.Println("Body:")
		//io.Copy(os.Stdout, r.Body)
	})

	http.HandleFunc("/data/", func(writer http.ResponseWriter, request *http.Request) {
		// 对于自定义类型需要获取body原始数据
		// 使用特定解码器
		// io.Copy(os.Stdout, request.Body)
		ctx, _ := ioutil.ReadAll(request.Body)
		var j map[string]interface{}
		json.Unmarshal(ctx, &j)
		fmt.Printf("%#v\n", j)

		fmt.Fprint(writer, time.Now().Format("2020-09-09 12:09:02"))
	})
	http.ListenAndServe(addr, nil)
}




