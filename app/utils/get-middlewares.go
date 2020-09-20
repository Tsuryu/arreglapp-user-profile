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
		middlewares.ValidateToken,
	},
	"putUserProfile": []gin.HandlerFunc{
		middlewares.ValidateToken,
	},
	"deleteUserProfile": []gin.HandlerFunc{
		middlewares.ValidateToken,
	},
}

// GetMiddlewares : get array of middlewares by name
func GetMiddlewares(name string) []gin.HandlerFunc {
	return middlewarelist[name]
}
