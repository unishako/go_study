package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var instance DeptsHandler

type DeptsHandler interface {
	List(context *gin.Context)
}

type deptsHandler struct {
	message string
}

func NewDeptsHandler(msg string) DeptsHandler {
	// シングルトン対応（ここまでやるの？）
	if instance == nil {
		instance = deptsHandler{
			message: msg,
		}
	}
	return instance
}

func (handler deptsHandler) List(context *gin.Context) {
	logger, _ := zap.NewDevelopment()
	logger.Info(handler.message)
	return
}
