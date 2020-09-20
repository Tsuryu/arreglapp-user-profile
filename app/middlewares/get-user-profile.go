package middlewares

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/gin-gonic/gin"
)

// GetUserProfile : handle get user profile request/response
func GetUserProfile(context *gin.Context) {
	userProfile, err := userprofileservice.FindBy(context.Param("id"))
	if err != nil {
		context.JSON(404, gin.H{"message": "User profile " + context.Param("id") + " has not been found"})
		context.Abort()
		return
	}

	context.JSON(200, userProfile)
}
