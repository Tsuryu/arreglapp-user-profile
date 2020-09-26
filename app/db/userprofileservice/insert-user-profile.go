package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
)

// Insert : creates an user
func Insert(userProfile models.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// passwordEncrypted, err := utils.Encrypt(userProfile.Password)
	// if err != nil {
	// 	return err
	// }
	// userProfile.Password = passwordEncrypted
	userProfile.CreationDate = time.Now()
	userProfile.Status = "created"

	_, err := Collection.InsertOne(ctx, userProfile)
	return err
}
