package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// PostUserProfile : handle post user profile request/response
func PostUserProfile(context *gin.Context) {
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)

	err := userprofileservice.Insert(userProfile)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user", "user": userProfile})
		context.Abort()
		return
	}
}
