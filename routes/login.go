package routes

import (
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/config"
	"github.com/devfurkankizmaz/go-user-auth-service/controller"
	"github.com/devfurkankizmaz/go-user-auth-service/model"
	"github.com/devfurkankizmaz/go-user-auth-service/mongo"
	"github.com/devfurkankizmaz/go-user-auth-service/repository"
	"github.com/devfurkankizmaz/go-user-auth-service/service"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, model.CollectionUser)
	lc := &controller.LoginController{
		LoginService: service.NewLoginService(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
