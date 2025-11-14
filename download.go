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
	"sync"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

func downloadAll(urls []string, baseDir string) {
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	client := resty.New().SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))

	for i, u := range urls {
		sem <- struct{}{}
		wg.Add(1)

		go func(idx int, imgURL string) {
			defer wg.Done()
			defer func() { <-sem }()

			fileName := filepath.Join(baseDir, fmt.Sprintf("%03d.jpg", idx+1))
			resp, err := client.R().SetOutput(fileName).Get(imgURL)
			if err != nil || resp.IsError() {
				fmt.Printf("失败 %03d : %v / HTTP %d\n", idx+1, err, resp.StatusCode())
				return
			}
			fmt.Printf("完成 %03d → %s\n", idx+1, fileName)
		}(i, u)
	}

	wg.Wait()
	fmt.Println("全部下载任务结束！")
}
