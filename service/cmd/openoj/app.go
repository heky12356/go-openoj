package main

import (
	"go-openoj/service/internal/config"
	"go-openoj/service/internal/router"
)

func main() {
	r := router.Router()
	err := r.Run(":" + config.GetConfig().Server.Port)
	if err != nil {
		panic(err)
	}
}
