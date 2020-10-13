package routers

import (
	"os"

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
	router.POST("/user-profile/:id/reset-password", utils.GetMiddlewares("resetPassword")...)
	router.PUT("/user-profile/:id/password", utils.GetMiddlewares("updatePassword")...)
	router.Run(":" + os.Getenv("APP_PORT"))
}
