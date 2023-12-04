package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 1.将日志写入文件而不是终端
// 使用zap.New()来传递所有配置
// func New(core zapcore.Core, options ...Option) *Logger
// zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel
// Encoder:编码器(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()
// zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

// WriterSyncer ：指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去。

// file, _ := os.Create("./test.log")
// writeSyncer := zapcore.AddSync(file)

// Log Level：哪种级别的日志将被写入

var sugarLogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	// 如果希望将JsonEncoder 更改为普通的log Encoder
	// return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLoggerWrite() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func InitLogger() {
	writeSyncer := getLoggerWrite()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s ", url, err)
	} else {
		sugarLogger.Info("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
