package routers

import (
	utils "github.com/Tsuryu/arreglapp-user-profile/app/middlewares/utils"
	"github.com/gin-gonic/gin"
)

// Router : init paths handlers
func Router() {
	router := gin.Default()
	router.LoadHTMLGlob("./app/templates/*.html")
	router.POST("/login", utils.GetMiddlewares("login")...)
	router.POST("/user-profile", utils.GetMiddlewares("postUserProfile")...)
	router.PUT("/user-profile/:id/activate", utils.GetMiddlewares("activateProfile")...)
	router.Run()
}
