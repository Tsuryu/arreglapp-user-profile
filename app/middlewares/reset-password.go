package middlewares

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ResetPassword : create an otp to change password and updates user profile status
func ResetPassword(context *gin.Context) {
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)
	userProfile.Status = "resetPassword"

	userProfileDB, err := userprofileservice.FindByEmail(userProfile.Email)

	if err != nil {
		context.JSON(500, gin.H{"message": "Failed to reset password", "user": userProfile})
		context.Abort()
		return
	}

	keys := make(map[string]interface{})
	keys["user-profile"] = userProfileDB
	context.Keys = keys
}
