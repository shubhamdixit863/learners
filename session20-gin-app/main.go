package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// custom types in golang
func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := r.Run("localhost:8090")
	if err != nil {
		log.Println("Error starting the gin server")
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
