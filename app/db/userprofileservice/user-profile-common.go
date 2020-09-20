package userprofileservice

import "github.com/Tsuryu/arreglapp-user-profile/app/db"

var database = db.Connection.Database("arreglapp")

// Collection : collection database mongo name
var Collection = database.Collection("user_profile")
