package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)



func parse(line string) (string, string, int, error) {
	elements := strings.Fields(line)
	if len(elements) < 8 {
		return "", "", 0, fmt.Errorf("format error")
	}
	elements[0]
	elements[6]
	statusCode, err := strconv.Atoi(elements[8])
	if err != nil {
		return "", "", 0, err
	}

}

func main() {
	file, err := os.Open("ticket_access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	ipStats := make(map[string]int)

	for {
		ctx, _, err := reader.ReadLine()
		fmt.Println(err, string(ctx))
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if ip, url, statusCode err := parse(string(ctx)); err == nil {
			ipStats[ip]++
			urlStats[url]++
			statusCodeStats[statusCode]++
		}
	}
	fmt.Println(ipStats)
	fmt.Println(statusCodeStats)

}
