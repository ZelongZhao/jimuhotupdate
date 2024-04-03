package main

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/pkg/data"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/server"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/server/service"
	"github.com/google/wire"
)

var set = wire.NewSet(
	data.Set,
	server.Set,
)

type services struct {
	u *service.UserService
}