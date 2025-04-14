package handlers

import (
	"session-23-gin-jwt/internal/repository"
	"session-23-gin-jwt/internal/services"
)

type Handler struct {
	repo       repository.DbRepository
	jwtService *services.JWTService
}

func NewHandler(repo repository.DbRepository, jwtService *services.JWTService) *Handler {
	return &Handler{
		repo:       repo,
		jwtService: jwtService,
	}
}
