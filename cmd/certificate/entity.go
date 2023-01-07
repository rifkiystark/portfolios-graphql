package certificate

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certificate struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:",omitempty"`
	PublishedAt time.Time          `bson:"publishedAt,omitempty"`
	Description string             `bson:",omitempty"`
	Url         string             `bson:",omitempty"`
}

type CreateCertificateRequest struct {
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"publishedAt"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
}

type CertificateResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	PublishedAt time.Time `json:"publishedAt"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
}

type UpdateCertificateRequest struct {
	Title       *string    `json:"title"`
	PublishedAt *time.Time `json:"publishedAt"`
	Description *string    `json:"description"`
	Url         *string    `json:"url"`
}
