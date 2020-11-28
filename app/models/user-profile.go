package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserProfile : mongo db model, bson = database format data, json = response data
type UserProfile struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Username     string             `bson:"username" json:"username,omitempty"`
	Password     string             `bson:"password" json:"password,omitempty"`
	FirstName    string             `bson:"firstName" json:"first_name,omitempty"`
	LastName     string             `bson:"lastName" json:"last_name,omitempty"`
	CreationDate time.Time          `bson:"creationDate" json:"-"`
	Status       string             `json:"-"`
	Email        string             `bson:"email" json:"email,omitempty"`
	Phone        string
	City         string
	Address      string
	Token        string `bson:"token" json:"token,omitempty"`
}
