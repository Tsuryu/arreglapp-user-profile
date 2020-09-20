package middlewares

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
)

// PostUserProfile : handle post user profile request/response
func PostUserProfile(context *gin.Context) {
	userProfile := models.UserProfile{}
	context.ShouldBind(&userProfile)

	err := userprofileservice.Insert(userProfile)
	if err != nil {
		context.JSON(500, gin.H{"message": "Failed to create user", "user": userProfile})
		context.Abort()
		return
	}

	context.JSON(201, gin.H{"message": "ok"})
}
