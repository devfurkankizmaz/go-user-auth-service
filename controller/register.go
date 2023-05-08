package controller

import (
	"fmt"
	"net/http"

	"github.com/devfurkankizmaz/go-user-auth-service/config"
	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/devfurkankizmaz/go-user-auth-service/utils"
	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	RegisterService model.RegisterService
	Env             *config.Env
}

func (r *RegisterController) Register(c *gin.Context) {
	var payload model.RegisterRequest

	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	_, err = r.RegisterService.GetUserByEmail(c, payload.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"Message": "User already exists"})
		return
	}

	err = utils.EmailValid(payload.Email)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"Message": err.Error()})
		return
	}

	err = utils.VerifyPassword(payload.Password, payload.ConfirmPassword)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"Message": err.Error()})
		return
	}

	err = r.RegisterService.Save(c, &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message:": fmt.Sprintf("Inserted User Email: %v", payload.Email)})
}
