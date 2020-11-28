package middlewares

import (
	"net/http"

	"github.com/Tsuryu/arreglapp-user-profile/app/db/userprofileservice"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// LogIn : checks if it's a valid user
func LogIn(context *gin.Context) {
	login := models.Login{}
	context.ShouldBindBodyWith(&login, binding.JSON)

	userProfile, err := userprofileservice.FindBy(login.Username, login.Password)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if userProfile.Status != "active" {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if context.Keys != nil {
		context.Keys["userProfile"] = userProfile
	} else {
		keys := make(map[string]interface{})
		keys["userProfile"] = userProfile
		context.Keys = keys
	}
}
