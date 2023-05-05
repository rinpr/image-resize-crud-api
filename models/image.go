package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImageData struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Time string	`json:"time,omitempty" bson:"time,omitempty"`
	Path	string	`json:"path,omitempty" bson:"path,omitempty"`
	SizeBefore	string	`json:"sizebefore,omitempty" bson:"sizebefore,omitempty"`
	SizeAfter	string	`json:"sizeafter,omitempty" bson:"sizeafter,omitempty"`
	IsSuccess	bool	`json:"issuccess,omitempty" bson:"issuccess,omitempty"`
}