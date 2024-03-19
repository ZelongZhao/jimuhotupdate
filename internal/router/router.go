package router

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/api/v1"
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	authG := r.Group("/auth")
	{
		authG.POST("/login", v1.AuthLoginHandler)
	}

	g1 := r.Group("/v1")
	g1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware())
	{

		g1.GET("/hello", v1.AuthLoginHandler)
	}
	return r
}
