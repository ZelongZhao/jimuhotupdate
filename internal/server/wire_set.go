package server

import (
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server/repo"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server/service"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server/usecase"
	"github.com/google/wire"
)

// Set for di
var Set = wire.NewSet(
	service.NewUserService,
	usecase.NewUserUseCase,
	repo.NewUserRepo,
)
