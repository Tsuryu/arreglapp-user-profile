package middlewareutils

import (
	commonMiddlewares "github.com/Tsuryu/arreglapp-commons/app/middlewares"
	"github.com/Tsuryu/arreglapp-user-profile/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"postUserProfile": {
		middlewares.PostUserProfile,
		middlewares.CreateUserProfileTransaction,
	},
	"activateProfile": {
		middlewares.ValidateTransaction,
		middlewares.ActivateUserProfile,
		middlewares.FinishTransaction,
	},
	"login": {
		middlewares.LogIn,
		middlewares.UpdateDeviceToken,
		middlewares.CreateJWT,
	},
	"resetPassword": {
		middlewares.ResetPassword,
		middlewares.CreateResetPasswordTransaction,
	},
	"updatePassword": {
		middlewares.ValidateTransaction,
		middlewares.UpdatePassword,
		middlewares.FinishTransaction,
	},
	"pushNotification": {
		middlewares.GetPushNotificationID,
	},
	"getMyProfile": {
		commonMiddlewares.ValidateJwt,
		middlewares.GetMyProfile,
	},
	"putUserProfile": {
		commonMiddlewares.ValidateJwt,
		middlewares.PutUserProfile,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
