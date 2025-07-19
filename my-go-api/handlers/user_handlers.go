package handlers

import (
	"net/http" // Importe para usar http.Status*

	"my-go-api/models" // Importa o pacote models

	"github.com/gin-gonic/gin"
)

// users simula um "banco de dados" em memória
var users = []models.User{
	{Name: "Enzo Kasma", Email: "enzokasma@gmail.com"},
}

// ShowUsers retorna todos os usuários.
func ShowUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

// CreateUser cria um novo usuário.
func CreateUser(ctx *gin.Context) {
	newUser := models.User{}
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	users = append(users, newUser)
	ctx.JSON(http.StatusCreated, newUser)
}
