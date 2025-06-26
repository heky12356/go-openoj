package handler

import (
	"fmt"
	"log"
	"net/http"

	"go-judge/internal/service"

	"github.com/gin-gonic/gin"
)

type submit struct {
	Code  string `json:"code"`
	Stdin string `json:"stdin"`
}

func HandleSubmit(c *gin.Context) {
	var codedata submit
	err := c.ShouldBind(&codedata)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bind failed: %s", err)})
		return
	}
	log.Print(codedata)
	resault, err := service.ServiceSubmit(codedata.Code, codedata.Stdin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("service error: %s", err.Error())})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"output": resault})
}
