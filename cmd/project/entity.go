package project

import (
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Title          string             `bson:",omitempty"`
	ImageURL       string             `bson:"imageURL,omitempty"`
	ImageID        string             `bson:"imageID,omitempty"`
	AdditionalInfo []string           `bson:",omitempty"`
	Description    string             `bson:",omitempty"`
}

type CreateProjectRequest struct {
	Title          string         `json:"title"`
	Image          graphql.Upload `json:"image"`
	AdditionalInfo []string       `json:"additionalInfo"`
	Description    string         `json:"description"`
}

type ProjectResponse struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	ImageURL       string   `json:"imageUrl"`
	AdditionalInfo []string `json:"additionalInfo"`
	Description    string   `json:"description"`
}

type UpdateProjectRequest struct {
	Title          *string         `json:"title"`
	Image          *graphql.Upload `json:"image"`
	AdditionalInfo []string        `json:"additionalInfo"`
	Description    *string         `json:"description"`
}
