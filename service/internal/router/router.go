package router

import (
	"go-openoj/service/internal/handler"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	r.POST("/submit", handler.HandleSubmit)
	return r
}
