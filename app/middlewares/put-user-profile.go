package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
)

// PutUserProfile : udpates a user profile
func PutUserProfile(context *gin.Context) {
	claim := context.Keys["claims"].(*commonModels.Claim)
	userProfile := models.UserProfile{}
	context.ShouldBind(&userProfile)

	userProfile.Username = claim.Username
	modifiedCount, err := userprofileservice.Update(userProfile)
	if err != nil || modifiedCount == 0 {
		context.JSON(404, gin.H{"message": "User profile " + context.Param("id") + " has not been found"})
		context.Abort()
		return
	}

	userProfile.Password = ""

	context.JSON(http.StatusOK, userProfile)

}
