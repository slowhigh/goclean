package infra

import (
	"github.com/google/wire"
	"github.com/slowhigh/goclean/infra/config"
	"github.com/slowhigh/goclean/infra/database"
	"github.com/slowhigh/goclean/infra/database/repository"
	"github.com/slowhigh/goclean/infra/router"
)

var InfraSet = wire.NewSet(
	config.NewConfig,
	router.NewRouter,
	database.NewDatabase,
	repository.NewMemo,
)
