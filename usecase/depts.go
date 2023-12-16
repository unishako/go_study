package usecase

import (
	"go.uber.org/zap"
)

var instance DeptsUsecase

type DeptsUsecase interface {
	getEmps()
}

type deptsUsecase struct {
	message string
}

func NewDeptsUsecase(msg string) DeptsUsecase {
	// シングルトン対応（ここまでやるの？）
	if instance == nil {
		instance = deptsUsecase{
			message: msg,
		}
	}
	return instance
}

func (handler deptsUsecase) getEmps() {
	logger, _ := zap.NewDevelopment()
	logger.Info(handler.message)
	return
}
