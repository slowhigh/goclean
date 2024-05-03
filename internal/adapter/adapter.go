package adapter

import (
	"github.com/google/wire"
	"github.com/slowhigh/goclean/internal/adapter/controller/rest"
)

var AdapterSet = wire.NewSet(rest.NewMemoController)
