package utils

import "golang.org/x/crypto/bcrypt" // Importa o pacote bcrypt para hashing de senhas

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byte), err
} // Gera um hash da senha usando bcrypt

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // Compara a senha fornecida com o hash armazenado
}
