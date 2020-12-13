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
	elements := strings.Fields(line) // 返回一个切片
	if len(elements) < 8 {
		return "", "", 0, fmt.Errorf("format error")
	}
	statusCode, err := strconv.Atoi(elements[8])
	if err != nil {
		return "", "", 0, err
	}
	return elements[0], elements[6], statusCode, nil
}
func main() {
	// 1. 打开日志文件
	// 2. 从日志中每行取出IP
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	ipStats := make(map[string]int)
	urlStats := make(map[string]int)
	statusCodeStats := make(map[int]int)
	for {
		ctx, _, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if ip, url, statusCode, err := parse(string(ctx)); err == nil {
			ipStats[ip]++
			urlStats[url]++
			statusCodeStats[statusCode]++
		}
	}
	fmt.Println(ipStats)
	fmt.Println(urlStats)
	fmt.Println(statusCodeStats)
	ipStatsSlice := make([][]interface{}, 0, len(ipStats))
	for k, v := range ipStats {
		ipStatsSlice = append(ipStatsSlice, []interface{}{k, v})
	}
	fmt.Println(ipStatsSlice)
}