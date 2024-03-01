package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func greeting(c *gin.Context) {
	c.String(http.StatusOK, "Hello Post")
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	r.POST("/testPost", greeting)
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		newAction := strings.Trim(action, "/")
		message := name + "猛" + newAction

		c.String(http.StatusOK, message)

	})
	r.Run(":8080")

}
