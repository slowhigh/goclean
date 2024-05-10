//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/slowhigh/goclean/infra"
	"github.com/slowhigh/goclean/infra/router"
	"github.com/slowhigh/goclean/internal/adapter"
	"github.com/slowhigh/goclean/internal/usecase"
)

func InitServer() (router.Router, error) {
	wire.Build(infra.InfraSet, adapter.AdapterSet, usecase.UsecaseSet)
	return router.Router{}, nil
}
