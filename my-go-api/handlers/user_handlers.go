package handlers

import (
	"my-go-api/models"
	"my-go-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var users = []models.User{}

var nextUserID int = 1

func init() {
	hashedPasswordEnzo, _ := utils.HashPassword("senha123")
	hashedPasswordAdmin, _ := utils.HashPassword("admin123") // Senha para o admin

	users = []models.User{
		{ID: generateUserID(), Name: "Enzo Kasma", Email: "enzokasma@gmail.com", Password: hashedPasswordEnzo, Role: "user"},
		{ID: generateUserID(), Name: "Admin User", Email: "admin@example.com", Password: hashedPasswordAdmin, Role: "admin"},
	}
}

func generateUserID() string {
	id := strconv.Itoa(nextUserID)
	nextUserID++
	return id
}

// Cria uma instância do validador
var validate = validator.New()

func ShowUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

func CreateUser(ctx *gin.Context) {
	newUser := models.User{}
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload", "details": err.Error()})
		return
	}

	// Executa a validação na struct newUser
	if err := validate.Struct(newUser); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorsResponse := make(map[string]string)
		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "required":
				errorsResponse[fieldErr.Field()] = fieldErr.Field() + " é obrigatório."
			case "min":
				errorsResponse[fieldErr.Field()] = fieldErr.Field() + " deve ter no mínimo " + fieldErr.Param() + " caracteres."
			case "max":
				errorsResponse[fieldErr.Field()] = fieldErr.Field() + " deve ter no máximo " + fieldErr.Param() + " caracteres."
			case "email":
				errorsResponse[fieldErr.Field()] = fieldErr.Field() + " deve ser um endereço de e-mail válido."
			default:
				errorsResponse[fieldErr.Field()] = "Erro de validação para " + fieldErr.Field()
			}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errorsResponse})
		return
	}

	// Verifica se o email já existe
	for _, user := range users {
		if user.Email == newUser.Email {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
	}

	users = append(users, newUser)
	ctx.JSON(http.StatusCreated, newUser)
}
