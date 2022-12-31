package graph

import (
	"github.com/imagekit-developer/imagekit-go"
	"github.com/rifkiystark/portfolios-api/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *database.DB
	IK *imagekit.ImageKit
}
