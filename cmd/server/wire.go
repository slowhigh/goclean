//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/slowhigh/goclean/infra"
	"github.com/slowhigh/goclean/infra/router"
	"github.com/slowhigh/goclean/internal/controller"
	"github.com/slowhigh/goclean/internal/usecase"
)

func InitServer() (router.Router, error) {
	wire.Build(infra.InfraSet, controller.ControllerSet, usecase.UsecaseSet)
	return router.Router{}, nil
}
