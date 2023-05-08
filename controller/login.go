package controller

import (
	"net/http"

	"github.com/devfurkankizmaz/go-user-auth-service/config"
	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginService model.LoginService
	Env          *config.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var payload model.LoginRequest

	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	user, err := lc.LoginService.GetUserByEmail(c, payload.Email)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"Message": "User not found"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Invalid password"})
		return
	}

	at, err := lc.LoginService.GenAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	rt, err := lc.LoginService.GenRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": at, "refresh_token": rt})

}
