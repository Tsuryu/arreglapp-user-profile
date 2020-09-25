package utils

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/middlewares"
	"github.com/gin-gonic/gin"
)

var middlewarelist = map[string][]gin.HandlerFunc{
	"getUserProfile": []gin.HandlerFunc{
		middlewares.GetUserProfile,
	},
	"postUserProfile": []gin.HandlerFunc{
		middlewares.PostUserProfile,
		middlewares.SendConfirmationEmail,
	},
	"activateProfile": []gin.HandlerFunc{
		middlewares.ActivateUserProfile,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
