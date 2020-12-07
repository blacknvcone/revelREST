package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID        bson.ObjectId `json:"_id,omitempty", bson:"_id,omitempty"`
	FirstName string        `json: "first_name", bson:"first_name"`
	LastName  string        `json: "last_name", bson:"last_name"`
	Username  string        `json: "username", bson:"username"`
	Password  string        `json: "password", bson:"password"`
	Email     string        `json: "email", bson:"email"`
	CreatedAt string        `json: "created_at", bson:"created_at"`
	UpdatedAt string        `json: "Updated_at", bson::updated_at"`
}
