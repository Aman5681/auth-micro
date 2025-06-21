package routes

import (
	"net/http"

	"github.com/Aman5681/auth-micro/models"
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
	} else {
		context.JSON(http.StatusCreated, gin.H{"status": "success", "userId": userId})
	}

}
