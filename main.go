package main

import (
	"git.vfeda.com/vfeda/JiMuHotUpdate/internal/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":8080")
}
