package middlewares

import (
	"net/http"
	"os"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// ValidateTransaction : check if transaction has an confirmed detail step
func ValidateTransaction(context *gin.Context) {
	transaction := models.Transaction{}
	client := resty.New()

	request := client.R().SetResult(&transaction)

	traceID := context.GetHeader("trace-id")
	_, err := request.Get(os.Getenv("API_TRANSACTION_BASE_URL") + "/transaction/" + traceID)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Invalid transaction"})
		return
	}

	var validTransaction bool
	for _, detail := range transaction.Details {
		if detail.Status == "confirm" {
			validTransaction = true
		}
	}

	if !validTransaction {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Missing otp validation"})
		return
	}
}
