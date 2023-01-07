package ipr

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPR struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:",omitempty"`
	PublishedAt time.Time          `bson:"publishedAt,omitempty"`
	Description string             `bson:",omitempty"`
	Url         string             `bson:",omitempty"`
}

type CreateIPRRequest struct {
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"publishedAt"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
}

type IPRResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"publishedAt"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
}

type UpdateIPRRequest struct {
	Title       *string    `json:"title"`
	PublishedAt *time.Time `json:"publishedAt"`
	Description *string    `json:"description"`
	Url         *string    `json:"url"`
}
