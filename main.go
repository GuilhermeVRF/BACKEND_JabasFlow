package main

import (
	"jabas-flow/controllers"
	"jabas-flow/services"
	"jabas-flow/database"
	"jabas-flow/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dataBaseConnection, databaseError  := database.Connect()
	if(databaseError != nil){
		panic(databaseError)
	}

	UserRepository := repository.NewUserRepository(dataBaseConnection)
	UserService := services.NewUserService(UserRepository)
	UserController := controllers.NewUserController(UserService)

	// User routes
	server.GET("/users", UserController.Index)
	server.GET("/user/:id", UserController.Show)
	server.POST("/user", UserController.Store)
	server.PUT("/user/:id", UserController.Update)
	server.Run(":5000")
}
