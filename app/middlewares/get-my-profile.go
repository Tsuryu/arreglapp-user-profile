package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/gin-gonic/gin"
)

// GetMyProfile : handle get my profile
func GetMyProfile(context *gin.Context) {
	claim := context.Keys["claims"].(*commonModels.Claim)

	userProfile, err := userprofileservice.FindMyProfile(claim.Username)
	if err != nil {
		context.JSON(404, gin.H{"message": "User profile " + context.Param("id") + " has not been found"})
		context.Abort()
		return
	}

	userProfile.Password = ""

	context.JSON(http.StatusOK, userProfile)
}
