package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/jwt"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
)

// CreateJWT : creates a jwt to handle user's requests
func CreateJWT(context *gin.Context) {
	userProfile := context.Keys["userProfile"].(*models.UserProfile)
	jwtString, err := jwt.CreateJWT(userProfile)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create jwt"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwtString})
}
