package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Handler struct {
	users []User // dependency injection
}

func NewHandler(users []User) *Handler {
	return &Handler{users: users}
}

func (h *Handler) Create(c *gin.Context) {
	// give me the data

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(errors.New("Error in reading the data"))
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		c.Error(errors.New("Error creating the uui"))
		return
	}

	user.ID = id.String()

	// once you have the data ,you can insert in the
	h.users = append(h.users, user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfull",
		"Id":      id,
	})
}

func (h *Handler) Get(c *gin.Context) {
	// give me the data
	users := h.users

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"users":   users,
	})
}

func (h *Handler) GetById(c *gin.Context) {
	// give me the data
	//
	id := c.Param("id")

	for i := 0; i < len(h.users); i++ {
		if h.users[i].ID == id {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"user":    h.users[i],
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "No user Found",
		"users":   nil,
	})

}

func (h *Handler) Update(c *gin.Context) {
	// give me the data
	//
	id := c.Param("id")

	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.Error(errors.New("Error in reading the data"))
		return
	}
	user.ID = id

	for i := 0; i < len(h.users); i++ {
		if h.users[i].ID == id {

			h.users[i] = user
			c.JSON(http.StatusOK, gin.H{
				"message": "Update Success",
				"user":    h.users[i],
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "No user Found",
		"users":   nil,
	})

}
