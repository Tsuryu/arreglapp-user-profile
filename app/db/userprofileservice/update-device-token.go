package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateDeviceToken : updates data in mongo
func UpdateDeviceToken(userProfile models.UserProfile) error {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// map string, interface
	register := make(map[string]interface{})
	register["token"] = userProfile.Token

	updateString := bson.M{
		"$set": register,
	}

	filter := bson.M{
		"username": bson.M{
			"$eq": userProfile.Username, // gt
		},
		"password": bson.M{
			"$eq": userProfile.Password, // gt
		},
	}

	_, err := Collection.UpdateOne(context, filter, updateString)
	if err != nil {
		return err
	}

	return nil
}
