package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UserProfile : mongo db model, bson = database format data, json = response data
type UserProfile struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	FirstName string             `bson:"firstName" json:"first_name,omitempty"`
	LastName  string             `bson:"lastName" json:"last_name,omitempty"`
}
