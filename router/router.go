package router

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/Middlewares"
	"git.vfeda.com/vfeda/JiMuHotUpdate/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	authG := r.Group("/auth")
	{
		authG.POST("/login", v1.AuthLoginHandler)
	}

	g1 := r.Group("/v1")
	g1.Use(Middlewares.JWTAuthMiddleware(), Middlewares.RateLimitMiddleware())
	{

		g1.GET("/hello", v1.AuthLoginHandler)
	}
	return r
}
