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
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

var cli = resty.New().SetTimeout(2 * time.Second)

func Bing_search(key_word string, count int) {
	safeDir := regexp.MustCompile(`[\\/:*?"<>|]`).ReplaceAllString(key_word, "_")
	_ = os.MkdirAll(safeDir, 0755)

	var urls []string
	page := 0

	for len(urls) < count {
		params := url.Values{}
		params.Add("q", key_word)
		params.Add("first", strconv.Itoa(page*35+1))
		params.Add("count", "35")
		searchURL := "https://cn.bing.com/images/search?" + params.Encode()

		resp, err := cli.R().Get(searchURL)
		if err != nil {
			fmt.Println("请求失败:", err)
			break
		}

		html := resp.String()
		reInner := regexp.MustCompile(`(?is)<div\b[^>]*\bimg_cont\b[^>]*\bhoff\b[^>]*>(.*?)</div>`)
		inners := reInner.FindAllStringSubmatch(html, -1)

		for _, m := range inners {
			m[1] = strings.ReplaceAll(m[1], "&amp;", "&")
			re := regexp.MustCompile(`\bsrc=["']([^"']+)["']`)
			if src := re.FindStringSubmatch(m[1]); len(src) > 1 {
				imgURL := src[1]
				if pos := strings.IndexByte(imgURL, '?'); pos != -1 {
					imgURL = imgURL[:pos]
				}
				urls = append(urls, imgURL)
				if len(urls) >= count {
					break
				}
			}
		}

		page++
		if page > 10 {
			break
		}
	}

	if len(urls) > count {
		urls = urls[:count]
	}

	fmt.Printf("共获取 %d 张图片，开始下载...\n", len(urls))
	downloadAll(urls, safeDir)
}
