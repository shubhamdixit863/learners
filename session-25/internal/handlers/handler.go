package handlers

import (
	"session-25/internal/repository"
)

type Handler struct {
	repo repository.Db
}
