package routes

import (
	"net/http"

	"github.com/Aman5681/auth-micro/models"
	"github.com/Aman5681/auth-micro/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind json"})
		return
	}
	userId, err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "success", "userId": userId})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind json"})
		return
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.EmailId, user.UserId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "could not authenticate user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "login success", "token": token})
}
