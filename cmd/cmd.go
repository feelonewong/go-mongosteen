package cmd

import (
	"go-mongosteen/internal/router"
)

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
