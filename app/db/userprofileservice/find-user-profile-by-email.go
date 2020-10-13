package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// FindByEmail : fetchs user profile by username
func FindByEmail(email string) (*models.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"email": email}

	var result *models.UserProfile

	err := Collection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
