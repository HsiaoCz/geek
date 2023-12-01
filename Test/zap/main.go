package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newCustomLogger() (*zap.Logger, error) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "", // 不记录日志调用位置
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "test.log"},
		ErrorOutputPaths: []string{"error.log"},
	}
	return cfg.Build()
}

func main() {
	// 生产环境
	{
		logger, _ := zap.NewProduction()
		defer logger.Sync() // 刷新buffer,保证日志最终会被输出

		url := "https://jianghushinian.cn"
		logger.Info("Production failed to fetch URL", zap.String("URL", url),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
	{
		// 开发环境
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()

		url := "https://jianghushinian.cn/"
		logger.Debug("Development failed to fetch URL",
			zap.String("URL", url),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
	{
		// 自定义logger
		logger, _ := newCustomLogger()
		defer logger.Sync()

		// 增加一个skip选项，触发zap内部error，将错误输出到error.log
		logger = logger.WithOptions(zap.AddCallerSkip(100))

		logger.Info("Info msg")
		logger.Error("Error msg")
	}
}
