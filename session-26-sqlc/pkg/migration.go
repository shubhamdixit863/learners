package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"session-23-gin-jwt/internal/models"
)

func Migrate() {
	// migration of tables

	dsn := "root:root@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err = db.AutoMigrate(&models.Products{})
	if err != nil {
		log.Println(err)
	}

	log.Println("Migration done !!!!")

}

func main() {
	Migrate()
}
