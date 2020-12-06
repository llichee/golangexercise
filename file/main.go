package main

import (
	"fmt"
	"os"
)

func main(){
	path := "testdir"  // 假设是一个目录
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// content := make([]byte, 10)
	// fmt.Println(file.Read(content))  // 不能直接读取目录内容
	fileInfo, err := file.Stat()
	fmt.Println(fileInfo, err)
}
