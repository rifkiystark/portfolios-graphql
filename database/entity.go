package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string             `bson:",omitempty"`
	ImageURL       string             `bson:",omitempty"`
	AdditionalInfo []string           `bson:",omitempty"`
	Description    string             `bson:",omitempty"`
}
