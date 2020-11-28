package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/gin-gonic/gin"
)

// GetPushNotificationID : get push notification id from an user
func GetPushNotificationID(context *gin.Context) {
	id, err := userprofileservice.FindUserProfileDeviceIDBy(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "User profile " + context.Param("id") + " has not been found"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, struct {
		ID string `json:"id"`
	}{id})
}
