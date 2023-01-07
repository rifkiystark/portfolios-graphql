package project

import (
	"context"
	"fmt"

	"github.com/rifkiystark/portfolios-api/internal/imagekit"
	"github.com/sirupsen/logrus"
)

type ProjectService interface {
	CreateProject(ctx context.Context, project CreateProjectRequest) (*ProjectResponse, error)
	UpdateProject(ctx context.Context, id string, project UpdateProjectRequest) (*ProjectResponse, error)
	DeleteProject(ctx context.Context, id string) (*ProjectResponse, error)
	GetProject(ctx context.Context, id string) (*ProjectResponse, error)
	GetProjects(ctx context.Context, search *string) ([]*ProjectResponse, error)
}

type ProjectServiceImpl struct {
	projectRepository ProjectRepository
	ik                imagekit.ImageKit
}

func NewService(projectRepository ProjectRepository, ik imagekit.ImageKit) ProjectService {
	return &ProjectServiceImpl{projectRepository: projectRepository, ik: ik}
}

func (p *ProjectServiceImpl) CreateProject(ctx context.Context, project CreateProjectRequest) (*ProjectResponse, error) {
	fileId, fileUrl, err := p.ik.Upload(ctx, project.Image.File, project.Image.Filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	projectEntity := project.ToEntity()
	projectEntity.ImageURL = fileUrl
	projectEntity.ImageID = fileId

	err = p.projectRepository.CreateProject(&projectEntity)
	if err != nil {
		p.ik.Delete(ctx, fileId)
		logrus.Errorf("error creating project: %v", err)
		return nil, err
	}

	projectModel := projectEntity.ToResponse()

	return &projectModel, nil
}

func (p *ProjectServiceImpl) UpdateProject(ctx context.Context, id string, updateProject UpdateProjectRequest) (*ProjectResponse, error) {
	if updateProject.AdditionalInfo == nil && updateProject.Description == nil && updateProject.Image == nil && updateProject.Title == nil {
		return nil, fmt.Errorf("no input provided")
	}

	projectEntity := updateProject.ToEntity()

	if updateProject.Image != nil {
		fileId, fileUrl, err := p.ik.Upload(ctx, updateProject.Image.File, updateProject.Image.Filename)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil, err
		}
		projectEntity.ImageURL = fileUrl
		projectEntity.ImageID = fileId
	}

	if err := p.projectRepository.UpdateProject(id, &projectEntity); err != nil {
		if updateProject.Image != nil {
			p.ik.Delete(ctx, projectEntity.ImageID)
		}
		return nil, err
	}

	projectModel := projectEntity.ToResponse()

	return &projectModel, nil
}

func (p *ProjectServiceImpl) DeleteProject(ctx context.Context, id string) (*ProjectResponse, error) {
	projectEntity, err := p.projectRepository.GetProject(id)
	if err != nil {
		return nil, err
	}

	logrus.Infof("project: %+v", projectEntity)
	err = p.projectRepository.DeleteProject(id)
	if err != nil {
		return nil, err
	}

	imageID := projectEntity.ImageID
	err = p.ik.Delete(ctx, imageID)
	if err != nil {
		logrus.Error(err)
	}

	projectModel := projectEntity.ToResponse()

	return &projectModel, nil
}

func (p *ProjectServiceImpl) GetProject(ctx context.Context, id string) (*ProjectResponse, error) {
	projectEntity, err := p.projectRepository.GetProject(id)
	if err != nil {
		return nil, err
	}

	projectModel := projectEntity.ToResponse()

	return &projectModel, nil
}

func (p *ProjectServiceImpl) GetProjects(ctx context.Context, search *string) ([]*ProjectResponse, error) {
	projectEntities, err := p.projectRepository.GetProjects(search)
	if err != nil {
		return nil, err
	}

	var projects []*ProjectResponse
	for _, projectEntity := range projectEntities {
		projectModel := projectEntity.ToResponse()
		projects = append(projects, &projectModel)
	}

	return projects, nil
}
