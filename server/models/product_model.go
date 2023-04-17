package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       float64            `bson:"price,omitempty"`
	Img         string             `bson:"img,omitempty"`
	InStock     bool               `bson:"inStock,omitempty"`
}
