package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateByEmail : updates data in mongo
func UpdateByEmail(userProfile models.UserProfile) (int64, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// map string, interface
	register := make(map[string]interface{})
	register["password"] = userProfile.Password

	updateString := bson.M{
		"$set": register,
	}

	filter := bson.M{
		"email": bson.M{
			"$eq": userProfile.Email, // gt
		},
	}

	result, err := Collection.UpdateOne(context, filter, updateString)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, err
}
