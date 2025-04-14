package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-mysql-org/go-mysql/driver"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"session-23-gin-jwt/internal/handlers"
	"session-23-gin-jwt/internal/middlewares"
	"session-23-gin-jwt/internal/repository"
	"session-23-gin-jwt/internal/services"
)

func main() {
	dsn := "root:root@localhost:3306?users"

	// We will load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find env file")
		return
	}

	r := gin.Default()

	r.Use(cors.Default())

	// Handler Object
	//repo := repository.NewInMemory()
	// sqlx connection
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to the db", err)
		return
	}
	repo := repository.NewMysqlReqo(db)
	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	v1.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "All good",
		})
	})
	// Two types of groups // auth routes
	auth := v1.Group("/auth") // /api/v1/auth

	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	// user routes
	user := v1.Group("/user") // /api/v1/user
	user.GET("/getUsers", middlewares.AuthorizationMiddleware(), handler.GetAllUsers)

	err = r.Run("localhost:8080")
	if err != nil {
		return
	}

}
