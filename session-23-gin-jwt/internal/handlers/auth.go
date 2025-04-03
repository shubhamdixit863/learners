package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	_, err = h.repo.GetUserByUserName(signupRequest.Username)
	log.Println(err)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User With Same Username already exists",
		})
		return
	}

	// What you have to do is save the data
	// We have to convert the dto to the model
	userIdData, _ := uuid.NewUUID()
	userModel := models.User{
		Username: signupRequest.Username,
		Password: signupRequest.Password,
		Name:     signupRequest.Name,
		Email:    signupRequest.Email,
		ID:       userIdData.String(),
	}

	userId := h.repo.CreateUser(userModel)
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
	user, err := h.repo.GetUserByUserName(signupRequest.Username)
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
