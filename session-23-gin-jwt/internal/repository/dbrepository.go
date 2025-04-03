package repository

import "session-23-gin-jwt/internal/models"

type DbRepository interface {
	CreateUser(user models.User) string
	GetUserByUserName(userName string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}
