package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {

	users, err := h.repo.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed To Get users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})

}
