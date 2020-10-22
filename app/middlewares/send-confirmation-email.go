package middlewares

import (
	utils "github.com/Tsuryu/arreglapp-commons/app/utils"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SendConfirmationEmail : handle confirmation email service call
func SendConfirmationEmail(context *gin.Context) {
	otp := context.Keys["otp"].(string)
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)

	go utils.SendEmail(userProfile.Email, otp)
}
