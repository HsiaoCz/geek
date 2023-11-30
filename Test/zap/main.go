package main

import (
	"time"

	"go.uber.org/zap"
)

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
}
