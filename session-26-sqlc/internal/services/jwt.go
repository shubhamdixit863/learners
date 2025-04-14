package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

type JWTService struct {
}

func (j *JWTService) CreateJWT(username string) (string, error) {
	// Create a new token with signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
	})

	secret := os.Getenv("secret")
	if secret == "" {
		return "", fmt.Errorf("JWT secret is not set")
	}

	// Sign the token with the secret
	return token.SignedString([]byte(secret))
}

func (j *JWTService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("secret")
	if secret == "" {
		return nil, fmt.Errorf("JWT secret is not set")
	}

	// Parse the token with a custom validation function
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is exactly HS256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["username"])
	} else {
		log.Printf("Invalid JWT Token")

	}

	if err != nil {
		log.Println("JWT validation error:", err)
		return nil, err
	}

	return token, nil
}
