package controller

import (
	"github.com/google/wire"
	"github.com/slowhigh/goclean/internal/controller/rest"
)

var ControllerSet = wire.NewSet(rest.NewMemoController)
