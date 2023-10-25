package moo

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 选项模式

type HTTPOption func(e *Engine)

// 这里添加一个优雅关机的选项

func WithHTTPStop(fn func() error) HTTPOption {
	return func(e *Engine) {
		if fn == nil {
			fn = func() error {
				fmt.Println("111111111")
				// os.Signal类型的channel
				quit := make(chan os.Signal)
				// 如果匹配到中断信息，会将信号传到channel里面
				// 如果没有匹配到，channel会一直阻塞
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
				log.Println("shutdown Server....")

				// 创建一个超时的上下文，在这里等五秒钟
				// 等待任务执行完毕
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := e.srv.Shutdown(ctx); err != nil {
					log.Fatal("server Shutdown", err)
				}
				// 关闭之后执行的操作
				select {
				case <-ctx.Done():
					log.Println("timeout of 5 seconds...")
				}
				return nil
			}
		}
		e.stop = fn
	}
}
