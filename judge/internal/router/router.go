package router

import (
	"go-judge/internal/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	return r
}
