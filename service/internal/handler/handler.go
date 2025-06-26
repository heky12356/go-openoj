package handler

import (
	"fmt"
	"net/http"

	"go-openoj/service/internal/define"
	"go-openoj/service/internal/service"

	"github.com/gin-gonic/gin"
)

func HandleSubmit(c *gin.Context) {
	var codedata define.Submit
	err := c.ShouldBind(&codedata)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("bind failed: %s", err)})
		return
	}
	resault, err := service.ServiceSubmit(codedata)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": resault})
}
