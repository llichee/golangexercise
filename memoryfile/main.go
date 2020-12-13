package main

import (
	"bytes"
	"fmt"
)

func main() {
	ctx := make([]byte, 5)

	buffer := bytes.NewBufferString("abc")
	buffer.WriteString("1234")
	n, err := buffer.Read(ctx)
	fmt.Println(err, string(ctx[:n]))
}
