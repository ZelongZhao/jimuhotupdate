package router

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/api/v1"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/pkg/middlewares"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/server/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(s service.UserService) *gin.Engine {
	r := gin.Default()

	authG := r.Group("/auth")
	{

		authG.POST("/login", s.AuthLoginHandler)
	}

	g1 := r.Group("/v1")
	g1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware())
	{

		g1.GET("/hello", v1.HelloHandler)
	}
	return r
}
