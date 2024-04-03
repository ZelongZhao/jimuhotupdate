package server

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/server/repo"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/server/service"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/server/usecase"
	"github.com/google/wire"
)

// Set for di
var Set = wire.NewSet(
	service.NewUserService,
	usecase.NewUserUseCase,
	repo.NewUserRepo,
)