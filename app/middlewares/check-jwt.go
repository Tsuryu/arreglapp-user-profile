package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/jwt"
	"github.com/gin-gonic/gin"
)

// CheckJWT : validates JWT
func CheckJWT(context *gin.Context) {
	auth := context.GetHeader("Authorization")
	err := jwt.ValidateJWT(auth)
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
