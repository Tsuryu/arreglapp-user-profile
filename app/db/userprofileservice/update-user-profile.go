package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Update : updates data in mongo
func Update(userProfile models.UserProfile) (int64, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// map string, interface
	register := make(map[string]interface{})
	if userProfile.Status != "" {
		register["status"] = userProfile.Status
	}
	updateString := bson.M{
		"$set": register,
	}

	filter := bson.M{
		"username": bson.M{
			"$eq": userProfile.Username, // gt
		},
	}

	result, err := Collection.UpdateOne(context, filter, updateString)

	return result.ModifiedCount, err
}
