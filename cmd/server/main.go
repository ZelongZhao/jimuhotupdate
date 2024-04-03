package main

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/pkg/router"
)



func main() {
	s,_ := NewServices()

	r := router.InitRouter(*s.u)

	r.Run(":8080")
}
