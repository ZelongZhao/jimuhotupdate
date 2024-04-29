package main

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/pkg/router"
)



func main() {
	s,_ := NewServices()

	r := router.NewRouter()

	r.Run(":8080")
}
