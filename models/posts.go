package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct{
	PID 			primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID				primitive.ObjectID `json:"UID" bson:"UID"`
	Caption 		string `json:"caption" bson:"caption"`
	ImageURL 		string `json:"image_url" bson:"image_url"`
	CreatedAt       string `json:"created_at,omitempty" bson:"created_at"`
}