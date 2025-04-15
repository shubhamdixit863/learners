package repository

import (
	"context"
	"session-23-gin-jwt/internal/models"
)

type DbRepository interface {
	CreateUser(ctx context.Context, user models.User) (interface{}, error)
	GetUserByUserName(ctx context.Context, userName string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, user models.User) error
}
