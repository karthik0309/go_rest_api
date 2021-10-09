package models

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	UID 			primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name 			string `json:"name" bson:"name"`
	Email 			string `json:"email" bson:"email"`
	HashedPassword 	string `json:"password,omitempty" bson:"hashed_password"`
}