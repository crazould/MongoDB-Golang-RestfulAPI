package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//User ...
type User struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NIM  string             `json:"nim,omitempty" bson:"NIM,omitempty"`
	Name string             `json:"name,omitempty" bson:"Name,omitempty"`
}
