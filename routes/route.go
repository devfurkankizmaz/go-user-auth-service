package routes

import (
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/config"
	"github.com/devfurkankizmaz/go-user-auth-service/mongo"
	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	rg := gin.Group("")
	NewRegisterRouter(env, timeout, db, rg)
}
