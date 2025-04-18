package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"session-23-gin-jwt/internal/dto"
	"session-23-gin-jwt/internal/models"
)

// https://stackoverflow.com/questions/3853749/what-is-the-difference-between-an-mvc-model-object-a-domain-object-and-a-dto
func (h *Handler) Signup(c *gin.Context) {

	var signupRequest dto.SignupRequest
	err := c.BindJSON(&signupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	// We check here if the user is already there
	_, err = h.repo.GetUserByUserName(c, signupRequest.Username)
	log.Println("Error in getting the user by username", err)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User With Same Username already exists",
		})
		return
	}

	// What you have to do is save the data
	// We have to convert the dto to the model

	userModel := models.User{
		Username:  signupRequest.Username,
		Password:  signupRequest.Password,
		FirstName: signupRequest.FirstName,
		SeconName: signupRequest.SecondName,
	}

	slog.Default().Debug("Data for the userModel", userModel)
	userId, err := h.repo.CreateUser(c, userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"ID":      userId,
		})

		return

	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"ID":      userId,
	})
	return

}

func (h *Handler) Login(c *gin.Context) {
	var signupRequest dto.LoginRequest
	err := c.BindJSON(&signupRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	// We check here if the user is already there
	user, err := h.repo.GetUserByUserName(c, signupRequest.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})

		return
	}

	// The jwt creation  code here

	jwt, err := h.jwtService.CreateJWT(user.Username)
	if err != nil {
		slog.Default().Error("Error in getting token", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed Generating the token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"Token":   jwt,
	})
	return

}
