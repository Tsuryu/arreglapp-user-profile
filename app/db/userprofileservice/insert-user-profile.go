package userprofileservice

import (
	"context"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"
	"golang.org/x/crypto/bcrypt"
)

// Insert : creates an user
func Insert(userProfile models.UserProfile) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bytes, _ := bcrypt.GenerateFromPassword([]byte(userProfile.Password), 6)
	userProfile.Password = string(bytes)
	userProfile.CreationDate = time.Now()
	userProfile.Status = "created"

	_, err := Collection.InsertOne(ctx, userProfile)
	return err
}
