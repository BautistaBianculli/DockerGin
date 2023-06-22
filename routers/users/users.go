package users

import (
	users "SatelitePacheco/database/users"
	"SatelitePacheco/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {

	users, err := users.GetUsers()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   users,
	})
}

func CreateUser(context *gin.Context) {
	var user models.Users

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}
	err := users.CreateUser(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "ERROR",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data": models.Users{
			Username: user.Username,
			Password: user.Password,
		},
		"Create Time": time.Now(),
	})
}
