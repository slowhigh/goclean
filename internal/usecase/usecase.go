package usecase

import (
	"github.com/google/wire"
	"github.com/slowhigh/goclean/internal/usecase/memo"
)

var UsecaseSet = wire.NewSet(memo.NewMemoUsecase)
