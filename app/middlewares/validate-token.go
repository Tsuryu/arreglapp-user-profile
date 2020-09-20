package middlewares

import (
	"github.com/gin-gonic/gin"
)

// ValidateToken : validate session token
func ValidateToken(context *gin.Context) {
	context.AbortWithStatus(501)
}
