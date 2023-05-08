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

func NewRegisterRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	r := repository.NewUserRepository(db, model.CollectionUser)
	c := controller.RegisterController{
		RegisterService: service.NewRegisterService(r, timeout),
		Env:             env,
	}
	group.POST("/register", c.Register)
}
