package ipr

import (
	"context"
	"fmt"

	"github.com/rifkiystark/portfolios-api/internal/imagekit"
	"github.com/sirupsen/logrus"
)

type IPRService interface {
	CreateIPR(ctx context.Context, ipr CreateIPRRequest) (*IPRResponse, error)
	UpdateIPR(ctx context.Context, id string, ipr UpdateIPRRequest) (*IPRResponse, error)
	DeleteIPR(ctx context.Context, id string) (*IPRResponse, error)
	GetIPR(ctx context.Context, id string) (*IPRResponse, error)
	GetIPRs(ctx context.Context, search *string) ([]*IPRResponse, error)
}

type IPRServiceImpl struct {
	iprRepository IPRRepository
	ik            imagekit.ImageKit
}

func NewService(iprRepository IPRRepository, ik imagekit.ImageKit) IPRService {
	return &IPRServiceImpl{iprRepository: iprRepository, ik: ik}
}

func (p *IPRServiceImpl) CreateIPR(ctx context.Context, ipr CreateIPRRequest) (*IPRResponse, error) {
	iprEntiry := ipr.ToEntity()

	err := p.iprRepository.CreateIPR(&iprEntiry)
	if err != nil {
		logrus.Errorf("error creating project: %v", err)
		return nil, err
	}

	iprResponse := iprEntiry.ToResponse()

	return &iprResponse, nil
}

func (p *IPRServiceImpl) UpdateIPR(ctx context.Context, id string, updateIPR UpdateIPRRequest) (*IPRResponse, error) {
	if updateIPR.Title == nil && updateIPR.Description == nil && updateIPR.PublishedAt == nil && updateIPR.Url == nil {
		return nil, fmt.Errorf("no input provided")
	}

	iprEntity := updateIPR.ToEntity()

	if err := p.iprRepository.UpdateIPR(id, &iprEntity); err != nil {
		return nil, err
	}

	iprResponse := iprEntity.ToResponse()

	return &iprResponse, nil
}

func (p *IPRServiceImpl) DeleteIPR(ctx context.Context, id string) (*IPRResponse, error) {
	iprEntity, err := p.iprRepository.GetIPR(id)
	if err != nil {
		return nil, err
	}

	logrus.Infof("ipr: %+v", iprEntity)
	err = p.iprRepository.DeleteIPR(id)
	if err != nil {
		return nil, err
	}

	iprResponse := iprEntity.ToResponse()

	return &iprResponse, nil
}

func (p *IPRServiceImpl) GetIPR(ctx context.Context, id string) (*IPRResponse, error) {
	projectEntity, err := p.iprRepository.GetIPR(id)
	if err != nil {
		return nil, err
	}

	projectModel := projectEntity.ToResponse()

	return &projectModel, nil
}

func (p *IPRServiceImpl) GetIPRs(ctx context.Context, search *string) ([]*IPRResponse, error) {
	iprEntities, err := p.iprRepository.GetIPRs(search)
	if err != nil {
		return nil, err
	}

	var iprs []*IPRResponse
	for _, projectEntity := range iprEntities {
		projectModel := projectEntity.ToResponse()
		iprs = append(iprs, &projectModel)
	}

	return iprs, nil
}
