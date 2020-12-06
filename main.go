package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("hello ni hao")
	connect := make([]byte, 6)
	file.Seek(0, 0)
	n, err := file.Read(connect)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	fmt.Println(connect)
}