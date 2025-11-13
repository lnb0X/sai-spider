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
