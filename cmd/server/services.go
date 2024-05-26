package main

import (
	v1 "git.vfeda.com/vfedabackendteam/jimuhotupdate/api/hotUpdate/v1"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/pkg/data"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var set = wire.NewSet(
	data.Set,
	server.Set,
)

type services struct {
	u *service.UserService
}

func (s *services) Register(r gin.IRouter) {
	v1.RegisterIMLoginServiceHTTPServer(r, s.u)
}
