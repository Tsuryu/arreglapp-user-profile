package middlewares

import (
	"fmt"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostUserProfile : handle post user profile request/response
func PostUserProfile(context *gin.Context) {
	fmt.Println("2")
	userProfile := models.UserProfile{}
	userProfile.ActivationCode = context.Keys["otp"].(string)
	context.ShouldBindBodyWith(&userProfile, binding.JSON)

	err := userprofileservice.Insert(userProfile)
	if err != nil {
		context.JSON(500, gin.H{"message": "Failed to create user", "user": userProfile})
		context.Abort()
		return
	}

	context.JSON(201, gin.H{"message": "ok"})
}
