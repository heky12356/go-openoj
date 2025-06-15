package main

import (
	"go-openoj/internal/config"
	"go-openoj/internal/router"
)

func main() {
	r := router.Router()
	err := r.Run(":" + config.GetConfig().Server.Port)
	if err != nil {
		panic(err)
	}
}
