package middlewares

import (
	"net/http"

	commonModels "github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/Tsuryu/arreglapp-commons/app/service"
	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateUserProfileTransaction : init transaction
func CreateUserProfileTransaction(context *gin.Context) {
	userProfile := models.UserProfile{}
	context.ShouldBindBodyWith(&userProfile, binding.JSON)

	transaction := commonModels.Transaction{}
	transactionDetail := commonModels.TransactionDetail{}

	transactionDetail.Metadata = struct {
		Step string `json:"step"`
	}{"init enrollment"}
	transactionDetail.Status = "created"
	transaction.Reference = "new-user"
	transaction.Email = userProfile.Email
	transaction.Details = append(transaction.Details, transactionDetail)

	err := service.NewTransaction(&transaction, true)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to init transaction"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"trace_id": transaction.TraceID})
}
