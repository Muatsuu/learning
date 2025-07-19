package main

import (
	"my-go-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Rotas para usuários
	router.GET("/users", handlers.ShowUsers)
	router.POST("/users", handlers.CreateUser)

	router.Run(":8080")
}
