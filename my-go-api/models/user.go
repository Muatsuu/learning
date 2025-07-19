package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Role     string `json:"role" validate:"required,oneof=admin user"` //Papel do usuario, user ou ADMIN
}
