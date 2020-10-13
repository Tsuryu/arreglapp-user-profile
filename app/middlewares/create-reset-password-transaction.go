package middlewares

import (
	"net/http"
	"os"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// CreateResetPasswordTransaction : init transaction
func CreateResetPasswordTransaction(context *gin.Context) {
	userProfile := context.Keys["user-profile"].(*models.UserProfile)
	transaction := models.Transaction{}
	client := resty.New()

	request := client.R()
	detail := `{"metadata":{"name": "` + userProfile.FirstName + `"}}`
	request.SetBody(`{"reference":"reset password", "email": "` + userProfile.Email + `", "details":[` + detail + `]}`)
	request.SetHeader("with-security-code", "true")
	request.SetResult(&transaction)

	_, err := request.Post(os.Getenv("API_TRANSACTION_BASE_URL") + "/transaction")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to init transaction"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"trace_id": transaction.TraceID})
}
