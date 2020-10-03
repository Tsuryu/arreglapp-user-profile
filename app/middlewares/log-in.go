package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
)

// LogIn : checks if it's a valid user
func LogIn(context *gin.Context) {
	login := models.Login{}
	context.ShouldBind(&login)

	userProfile, err := userprofileservice.FindBy(login.Username, login.Password)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if userProfile.Status != "active" {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	keys := make(map[string]interface{})
	keys["userProfile"] = userProfile
	context.Keys = keys
}
