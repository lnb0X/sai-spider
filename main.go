// SPDX-License-Identifier: GPL-3.0-or-later
// Copyright (C) 2025 xi~
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See

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
