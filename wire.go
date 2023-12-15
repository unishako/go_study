//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"golang/handler"
)

func Initialize(msg string) handler.DeptsHandler {
	wire.Build(handler.NewDeptsHandler)
	return nil
}
