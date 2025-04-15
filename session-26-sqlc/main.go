package main

import (
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/joho/godotenv"
	"log"
	"net/http"
	"session-23-gin-jwt/internal/handlers"
	"session-23-gin-jwt/internal/middlewares"
	"session-23-gin-jwt/internal/repository"
	"session-23-gin-jwt/internal/services"
)

func mongoConnect(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	//dsn := "root:root@localhost:3306?users"
	var dbtype string
	flag.StringVar(&dbtype, "dbtype", "mongodb", "type of database to connect to (e.g. mongodb, postgres)")
	flag.Parse()
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
	//db, err := sqlx.Connect("mysql", dsn)
	//if err != nil {
	//	log.Println("Error connecting to the db", err)
	//	return
	//}
	//repo := repository.NewMysqlReqo(db)
	//dsn := "root:root@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})
	//
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//repo := repository.NewMysqlOrm(db)
	var repo repository.DbRepository
	if dbtype == "mongodb" {
		client, err := mongoConnect("mongodb://localhost:27017/users")
		if err != nil {
			log.Fatal("Error connecting mongodb", err)
			return
		}
		repo = repository.NewMongoRepo(client)
	} else if dbtype == "mysql" {

	}

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
