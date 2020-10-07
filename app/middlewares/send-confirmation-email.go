package middlewares

import (
	"fmt"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/Tsuryu/arreglapp-user-profile/app/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// SendConfirmationEmail : handle confirmation email service call
func SendConfirmationEmail(context *gin.Context) {
	fmt.Println("5")
	otp := context.Keys["otp"].(string)
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)

	go services.SendOTPEmail(userProfile.Email, otp)
}
