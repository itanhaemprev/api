package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Categorie is a model for crud
type Categorie struct {
	ID         primitive.ObjectID `bson:"id" json:"id"`
	Title      string             `bson:"title" json:"title"`
	CreatedAt  primitive.DateTime `bson:"created_at" json:"created_at"`
	ModifiedAt primitive.DateTime `bson:"modified_at" json:"modified_at"`
}
