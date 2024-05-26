package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//s,_ := NewServices()

	//r := router.NewRouter()

	r := gin.Default()

	s, _ := NewServices()

	s.Register(r)

	//r.Run(":8080")

}
