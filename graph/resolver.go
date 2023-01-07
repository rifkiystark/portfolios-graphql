package graph

import (
	"github.com/rifkiystark/portfolios-api/cmd/project"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProjectService project.ProjectService
}
