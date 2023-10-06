package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// errorgruop
// 用来处理并发错误
// errorgroup其中有两个方法
// Go可以代替go关键字，开启一个协程，并且能够返回错误
// Wait()方法可以返回第一个非零的错误，并且阻塞直到go调用的方法都返回

func fetchUrl(urls []string) error {
	g := new(errgroup.Group)
	for _, url := range urls {
		// 启动groutine来进行方法
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println("方法成功,", url)
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Println("所有url均访问成功")
	return nil
}

func main() {
	urls := []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	err := fetchUrl(urls)
	if err != nil {
		log.Fatal(err)
	}
}
