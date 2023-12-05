package main

import (
	"io"
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
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLoggerWrite() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	// 多位置输出日志
	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}

func InitLogger() {
	writeSyncer := getLoggerWrite()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 添加将调用函数信息记录到日志中的功能
	// logger := zap.New(core, zap.AddCaller())
	// 当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

func Initlogger() *zap.Logger {
	encoder := getEncoder()
	// test.log 记录全量日志
	logF, _ := os.Create("./test.log")
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
	// test.err.log 记录Error级别的日志
	errF, _ := os.Create("./test.err.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)
	return zap.New(core, zap.AddCaller())
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
