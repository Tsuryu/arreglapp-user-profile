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

	err := userprofileservice.Update(userProfile)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to activate user", "id": userProfile.Username})
		return
	}

	context.HTML(http.StatusOK, "confirmation-email-result.html", gin.H{"username": userProfile.Username})
}
