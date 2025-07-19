package model

import (
	"first_server/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName string             `bson:"user_name" validate:"required"`
	Email    string             `bson:"email" validate:"required,email,unique"`
	Password string             `bson:"password" validate:"required"`
	Role     string             `bson:"role" default:"user" enum:"user,admin"`
}

func GetUserCollection() *mongo.Collection {
	if database.DB == nil {
		panic("Database not initialized")
	}

	return database.DB.Collection("users")
}
