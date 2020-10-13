package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdatePassword : updates an user profile password on reset password feature
func UpdatePassword(context *gin.Context) {
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)

	modifiedCount, err := userprofileservice.UpdateByEmail(userProfile)
	if modifiedCount == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}
