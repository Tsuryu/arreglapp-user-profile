package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindBy : fetchs user profile by username
func FindBy(username string) (*models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"username": username}

	var result *models.UserProfile

	err := Collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
