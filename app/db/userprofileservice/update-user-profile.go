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
	if userProfile.FirstName != "" {
		register["firstName"] = userProfile.FirstName
	}
	if userProfile.LastName != "" {
		register["lastName"] = userProfile.LastName
	}
	if userProfile.Phone != "" {
		register["phone"] = userProfile.Phone
	}
	if userProfile.Address != "" {
		register["address"] = userProfile.Address
	}
	if userProfile.Email != "" {
		register["email"] = userProfile.Email
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
