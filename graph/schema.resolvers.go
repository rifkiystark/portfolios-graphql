package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/rifkiystark/portfolios-api/cmd/certificate"
	"github.com/rifkiystark/portfolios-api/cmd/ipr"
	"github.com/rifkiystark/portfolios-api/cmd/project"
)

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input project.CreateProjectRequest) (*project.ProjectResponse, error) {
	return r.ProjectService.CreateProject(ctx, input)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, id string, input project.UpdateProjectRequest) (*project.ProjectResponse, error) {
	return r.ProjectService.UpdateProject(ctx, id, input)
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, id string) (*project.ProjectResponse, error) {
	return r.ProjectService.DeleteProject(ctx, id)
}

// CreateIPR is the resolver for the createIPR field.
func (r *mutationResolver) CreateIPR(ctx context.Context, input ipr.CreateIPRRequest) (*ipr.IPRResponse, error) {
	return r.IPRService.CreateIPR(ctx, input)
}

// UpdateIPR is the resolver for the updateIPR field.
func (r *mutationResolver) UpdateIPR(ctx context.Context, id string, input ipr.UpdateIPRRequest) (*ipr.IPRResponse, error) {
	return r.IPRService.UpdateIPR(ctx, id, input)
}

// DeleteIPR is the resolver for the deleteIPR field.
func (r *mutationResolver) DeleteIPR(ctx context.Context, id string) (*ipr.IPRResponse, error) {
	return r.IPRService.DeleteIPR(ctx, id)
}

// CreateCertificate is the resolver for the createCertificate field.
func (r *mutationResolver) CreateCertificate(ctx context.Context, input certificate.CreateCertificateRequest) (*certificate.CertificateResponse, error) {
	return r.CertificateService.CreateCertificate(ctx, input)
}

// UpdateCertificate is the resolver for the updateCertificate field.
func (r *mutationResolver) UpdateCertificate(ctx context.Context, id string, input certificate.UpdateCertificateRequest) (*certificate.CertificateResponse, error) {
	return r.CertificateService.UpdateCertificate(ctx, id, input)
}

// DeleteCertificate is the resolver for the deleteCertificate field.
func (r *mutationResolver) DeleteCertificate(ctx context.Context, id string) (*certificate.CertificateResponse, error) {
	return r.CertificateService.DeleteCertificate(ctx, id)
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, id string) (*project.ProjectResponse, error) {
	return r.ProjectService.GetProject(ctx, id)
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, search *string) ([]*project.ProjectResponse, error) {
	return r.ProjectService.GetProjects(ctx, search)
}

// Ipr is the resolver for the ipr field.
func (r *queryResolver) Ipr(ctx context.Context, id string) (*ipr.IPRResponse, error) {
	return r.IPRService.GetIPR(ctx, id)
}

// Iprs is the resolver for the iprs field.
func (r *queryResolver) Iprs(ctx context.Context, search *string) ([]*ipr.IPRResponse, error) {
	return r.IPRService.GetIPRs(ctx, search)
}

// Certificate is the resolver for the certificate field.
func (r *queryResolver) Certificate(ctx context.Context, id string) (*certificate.CertificateResponse, error) {
	return r.CertificateService.GetCertificate(ctx, id)
}

// Certificates is the resolver for the certificates field.
func (r *queryResolver) Certificates(ctx context.Context, search *string) ([]*certificate.CertificateResponse, error) {
	return r.CertificateService.GetCertificates(ctx, search)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
