package models

// Login : login struct
type Login struct {
	Username string `bson:"username" json:"username,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
}
