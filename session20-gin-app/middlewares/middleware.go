package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// we will access the info here

		// you can get headers easily
		auth := c.GetHeader("Authorization")

		if len(auth) == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Auth header missing",
			})
			c.Abort()
			return
		} else {
			log.Println("Middleware executed")
			c.Next()
		}
		return
	}
}
