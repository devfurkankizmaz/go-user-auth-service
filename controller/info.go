package controller

import (
	"net/http"

	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/gin-gonic/gin"
)

type InfoController struct {
	InfoService model.InfoService
}

func (ic *InfoController) Get(c *gin.Context) {
	userID := c.GetString("x-user-id")

	info, err := ic.InfoService.GetInfoByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, info)
}
