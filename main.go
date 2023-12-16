package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang/middleware"
)

func main() {

	// Loggerの設定
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			logger.Fatal("エラー", zap.Error(err))
		}
	}(logger)

	// FWの設定
	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger(logger))

	// ルーターの設定（router.goにした方がよい）
	deptsHandler := Initialize("通りました")

	v1 := engine.Group("/v1")
	depts := v1.Group("/depts")
	depts.GET("/emps", deptsHandler.List)
	if err := engine.Run(); err != nil {
		logger.Fatal("サーバが起動しませんでした。", zap.Error(err))
	}
}
