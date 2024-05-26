package router

import (
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/api/v1"
	"git.vfeda.com/vfedabackendteam/jimuhotupdate/internal/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewRouter,
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.RateLimitMiddleware())

	authG := r.Group("/auth")
	{
		authG.POST("/login")
	}

	g1 := r.Group("/v1")
	g1.Use(middlewares.JWTAuthMiddleware())
	{
		g1.GET("/hello", v1.HelloHandler)
	}
	return r
}
