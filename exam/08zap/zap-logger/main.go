package main

import (
	"net/http"

	"go.uber.org/zap"
)

// Zap提供了两种类型的日志记录器—Sugared Logger和Logger
// 在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录

// 在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录

// logger
// 通过调用zap.NewProduction()/zap.NewDevelopment()或者zap.Example()创建一个Logger
// 上面的每一个函数都将创建一个logger。唯一的区别在于它将记录的信息不同。例如production logger默认记录调用函数信息、日期和时间等。
// 通过Logger调用Info/Error等。
// 默认情况下日志都会打印到应用程序的console界面。

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info("Success ...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}

// 日志记录器方法的语法是这样的
// func (log *Logger) MethodXXX(msg string, fields ...Field)
// 其中MethodXXX是一个可变参数函数，可以是Info / Error/ Debug / Panic等。每个方法都接受一个消息字符串和任意数量的zapcore.Field场参数。
// 每个zapcore.Field其实就是一组键值对参数。