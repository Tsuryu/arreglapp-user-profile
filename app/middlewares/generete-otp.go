package middlewares

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

// GenerateOTP : creates an otp to use in different features
func GenerateOTP(context *gin.Context) {
	fmt.Println("1")
	random := fmt.Sprintf("%02d", rand.Intn(999999-0))

	keys := make(map[string]interface{})
	keys["otp"] = random
	context.Keys = keys
}
