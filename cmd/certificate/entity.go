package certificate

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certificate struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:",omitempty"`
	ValidUntil time.Time          `bson:"validUntil,omitempty"`
	Url        string             `bson:",omitempty"`
}

type CreateCertificateRequest struct {
	Title      string    `json:"title"`
	ValidUntil time.Time `json:"validUntil"`
	Url        string    `json:"url"`
}

type CertificateResponse struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	ValidUntil time.Time `json:"validUntil"`
	Url        string    `json:"url"`
}

type UpdateCertificateRequest struct {
	Title      *string    `json:"title"`
	ValidUntil *time.Time `json:"validUntil"`
	Url        *string    `json:"url"`
}
