package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()	{
	// 解析模板 本身就是字符串
	text := "my name is {{.}}"
	tpl, err := template.New("test").Parse(text)
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(os.Stdout, "xzw")
	tpl.Execute(os.Stdout, 123)
	tpl.Execute(os.Stdout, []int{5,7,9})

	text = "性别：{{ if . }} 男 {{ else }} 女 {{ end }} \n"
	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, false)

	text = "性别：{{ if eq 1 . }} 男 {{ else }} 女 {{ end }} \n"
	tpl, _ = template.New("test").Parse(text)
	tpl.Execute(os.Stdout, 1)
}
