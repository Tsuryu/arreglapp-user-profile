package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateDeviceToken : updates token to use through push notifications
func UpdateDeviceToken(context *gin.Context) {
	login := models.Login{}
	context.ShouldBindBodyWith(&login, binding.JSON)

	err := userprofileservice.UpdateDeviceToken(models.UserProfile{
		Username: login.Username,
		Password: login.Password,
		Token:    login.Token,
	})
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
