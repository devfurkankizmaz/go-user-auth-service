package main

import (
	"time"

	"github.com/devfurkankizmaz/go-user-auth-service/config"
	"github.com/devfurkankizmaz/go-user-auth-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := config.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()
	routes.Setup(env, timeout, db, gin)
	gin.Run(env.ServerAddress)
}
