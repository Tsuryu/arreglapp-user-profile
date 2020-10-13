package middlewareutils

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"postUserProfile": []gin.HandlerFunc{
		middlewares.GenerateOTP,
		middlewares.PostUserProfile,
		middlewares.SendConfirmationEmail,
	},
	"activateProfile": []gin.HandlerFunc{
		middlewares.ActivateUserProfile,
	},
	"login": []gin.HandlerFunc{
		middlewares.LogIn,
		middlewares.CreateJWT,
	},
	"resetPassword": []gin.HandlerFunc{
		middlewares.ResetPassword,
		middlewares.CreateResetPasswordTransaction,
	},
	"updatePassword": []gin.HandlerFunc{
		middlewares.ValidateTransaction,
		middlewares.UpdatePassword,
		middlewares.FinishTransaction,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
