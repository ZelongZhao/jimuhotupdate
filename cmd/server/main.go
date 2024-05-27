package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	s, _ := NewServices()

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()

	grpcClient, err := grpc.NewClient("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	gwMux := runtime.NewServeMux()
	r := gin.Default()

	s.Register(grpcServer, grpcClient, gwMux)

	r.Any("/*any", gin.WrapH(gwMux))

	err = r.Run(":1235")
	if err != nil {
		log.Fatal(err)
	}
}
