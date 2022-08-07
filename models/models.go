package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Body   string             `json:"body,omitempty" bson:"body,omitempty"`
	Author string             `json:"author,omitempty" bson:"author,omitempty"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Date     primitive.DateTime `json:"date,omitempty" bson:"date,omitempty"`
}
