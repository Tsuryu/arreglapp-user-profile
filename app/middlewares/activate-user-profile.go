package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
)

// ActivateUserProfile : change an user profile to active
func ActivateUserProfile(context *gin.Context) {
	userProfile := models.UserProfile{}
	userProfile.Username = context.Param("id")
	userProfile.Status = "active"

	modifiedCount, err := userprofileservice.Update(userProfile)
	if modifiedCount == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Failed to activate user", "id": userProfile.Username})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to activate user", "id": userProfile.Username})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}
