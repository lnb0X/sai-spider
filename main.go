package main

import (
    "os"
	"fmt"
	"math"
	"bufio"
	"strconv"
	"strings"
)

const maxConcurrency = 8

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("输入搜索关键词: ")
    key_word, _ := reader.ReadString('\n')
	key_word = strings.TrimSpace(key_word)

	fmt.Print("请输入下载数量: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	count := clampCount(line)

	fmt.Printf("准备下载 %d 张图\n", count)

	Bing_search(key_word, count)
}

func clampCount(line string) int {
	f, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 150
	}

	count := int(math.Ceil(f))

	switch {
	case count < 100:
		return 100
	case count > 300:
		return 300
	default:
		return count
	}
}
