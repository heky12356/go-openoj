package main

import (
	"go-judge/internal/router"
)

func main() {
	r := router.Router()
	err := r.Run(":5050")
	if err != nil {
		panic(err)
	}
}
