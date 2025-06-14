package handler

import (
	"fmt"
	"log"
	"net/http"

	"go-openoj/internal/model"
	"go-openoj/internal/service"

	"github.com/gin-gonic/gin"
)

func HandleSubmit(c *gin.Context) {
	var codedata model.Submit
	err := c.ShouldBind(&codedata)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bind failed: %s", err)})
		return
	}
	log.Print(codedata)
	resault, err := service.ServiceSubmit(codedata)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": resault})
}
