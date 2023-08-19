package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductModel struct {
	ProductId primitive.ObjectID `bson:"_id" json:"productId"`
	Name      string             `bson:"name" json:"name"`
	Price     int32              `bson:"price" json:"price"`
	Stock     int32              `bson:"stock" json:"stock"`
}
