package main

import (
	"context"
	v1 "git.vfeda.com/vfedabackendteam/jimuhotupdate/api/hotUpdate/v1"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/pkg/data"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/server/service"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var set = wire.NewSet(
	data.Set,
	server.Set,
)

type services struct {
	u *service.UserService
}

func (s *services) Register(grpcServer *grpc.Server, grpcLocalClient *grpc.ClientConn, restfulMux *runtime.ServeMux, authMux *runtime.ServeMux) {
	v1.RegisterIMLoginServiceServer(grpcServer, s.u)
	_ = v1.RegisterIMLoginServiceHandler(context.Background(), authMux, grpcLocalClient)
}
