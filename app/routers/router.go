package routers

import (
	"os"
	"strconv"

	utils "github.com/Tsuryu/arreglapp-user-profile/app/middlewares/utils"
	"github.com/gin-gonic/gin"
)

// Router : init paths handlers
func Router() {
	router := gin.Default()
	router.LoadHTMLGlob("./app/templates/*.html")
	router.POST("/session/login", utils.GetMiddlewares("login")...)
	router.POST("/user-profile", utils.GetMiddlewares("postUserProfile")...)
	router.GET("/my-profile", utils.GetMiddlewares("getMyProfile")...)
	router.PUT("/user-profile", utils.GetMiddlewares("putUserProfile")...)
	router.PUT("/user-profile/:id/activate", utils.GetMiddlewares("activateProfile")...)
	router.POST("/user-profile/:id/reset-password", utils.GetMiddlewares("resetPassword")...)
	router.PUT("/user-profile/:id/password", utils.GetMiddlewares("updatePassword")...)
	router.GET("/user-profile/:id/push-notification", utils.GetMiddlewares("pushNotification")...)

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err == nil {
		router.Run(":" + strconv.Itoa(port))
	} else {
		router.Run()
	}
}
