package repository

import "session-25/internal/models"

type Db interface {
	Insert(user models.User) error
}
