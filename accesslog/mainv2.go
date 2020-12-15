package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Parse(line string) (ip, url, statusCode string, err error) {
	strings.Fields(line)
}
func main() {
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	readers := bufio.NewReader(file)
	for {
		line, _, err := readers.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if ip, url, statusCode, err := Parse(line), err ==nil {

		}

	}
}
