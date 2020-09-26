package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindBy : fetchs user profile by username
func FindBy(username string, password string) (*models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"username": username, "password": password}

	var result *models.UserProfile

	err := Collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return nil, err
	}

	// data, err := utils.Decrypt(result.Password)
	// fmt.Println(data)
	// if err != nil {
	// 	return nil, err
	// }

	return result, nil
}
