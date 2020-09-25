package routers

import (
	"github.com/Tsuryu/arreglapp-user-profile/app/utils"
	"github.com/gin-gonic/gin"
)

// Router : init paths handlers
func Router() {
	router := gin.Default()
	router.LoadHTMLGlob("./app/templates/*.html")
	router.POST("/user-profile", utils.GetMiddlewares("postUserProfile")...)
	router.GET("/user-profile/:id", utils.GetMiddlewares("getUserProfile")...)
	router.GET("/user-profile/:id/activate", utils.GetMiddlewares("activateProfile")...)
	router.Run()
}
