package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindUserProfileDeviceIDBy : fetchs user profile and returns device id
func FindUserProfileDeviceIDBy(username string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"username": username}

	var result models.UserProfile

	err := Collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return "", err
	}

	// data, err := utils.Decrypt(result.Password)
	// fmt.Println(data)
	// if err != nil {
	// 	return nil, err
	// }

	return result.Token, nil
}
