package middlewares

import (
	"net/http"
	"os"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-resty/resty/v2"
)

// CreateUserProfileTransaction : init transaction
func CreateUserProfileTransaction(context *gin.Context) {
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)
	transaction := models.Transaction{}
	client := resty.New()

	request := client.R()
	detail := `{"metadata":{"step": "init enrollment"}}`
	request.SetBody(`{"reference":"create user profile", "email": "` + userProfile.Email + `", "details":[` + detail + `]}`)
	request.SetHeader("with-security-code", "true")
	request.SetResult(&transaction)

	_, err := request.Post(os.Getenv("API_TRANSACTION_BASE_URL") + "/transaction")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to init transaction"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"trace_id": transaction.TraceID})
}
