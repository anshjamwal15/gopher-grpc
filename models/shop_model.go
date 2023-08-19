package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type ShopModel struct {
	ShopId   primitive.ObjectID `bson:"_id" json:"shopId"`
	Name     string             `bson:"name" json:"name"`
	Users    []UserModel        `bson:"users" json:"users"`
	Location Location           `bson:"location" json:"location"`
}

func (m *ShopModel) Id() primitive.ObjectID {
	return m.ShopId
}
