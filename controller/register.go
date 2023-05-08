package controller

import (
	"fmt"
	"net/http"

	"github.com/devfurkankizmaz/go-user-auth-service/config"
	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/devfurkankizmaz/go-user-auth-service/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	hashedPassword, err := utils.HashPassword(payload.Password)
	payload.Password = hashedPassword
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	newUser := model.User{
		ID:       primitive.NewObjectID(),
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}

	err = r.RegisterService.Save(c, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message:": fmt.Sprintf("Inserted ID %v", newUser.ID.Hex())})
}
