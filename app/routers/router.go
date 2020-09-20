package routers

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/utils"
	"github.com/gin-gonic/gin"
)

// Router : init paths handlers
func Router() {
	router := gin.Default()
	router.POST("/user-profile", utils.GetMiddlewares("postUserProfile")...)
	router.GET("/user-profile/:id", utils.GetMiddlewares("getUserProfile")...)
	router.PUT("/user-profile/:id", utils.GetMiddlewares("putUserProfile")...)
	router.DELETE("/user-profile/:id", utils.GetMiddlewares("deleteUserProfile")...)
	router.Run()
}
