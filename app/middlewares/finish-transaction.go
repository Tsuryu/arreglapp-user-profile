package middlewares

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// FinishTransaction : close transaction so as to not be able to keep using it
func FinishTransaction(context *gin.Context) {
	client := resty.New()

	code := context.GetHeader("security-code")

	request := client.R()
	request.SetBody(`{"step":"finished"}`)
	request.SetHeader("security-code", code)
	request.SetHeader("last-update", "true")

	traceID := context.GetHeader("trace-id")

	_, err := request.Post(os.Getenv("API_TRANSACTION_BASE_URL") + "/transaction/" + traceID + "/detail")
	if err != nil {
		log.Println(err)
	}
}
