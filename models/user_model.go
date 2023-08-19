package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	UserId primitive.ObjectID `bson:"_id" json:"userId"`
	Name   string             `bson:"name" json:"name"`
}
