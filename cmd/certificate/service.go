package certificate

import (
	"context"
	"fmt"

	"github.com/rifkiystark/portfolios-api/internal/imagekit"
	"github.com/sirupsen/logrus"
)

type CertificateService interface {
	CreateCertificate(ctx context.Context, certificate CreateCertificateRequest) (*CertificateResponse, error)
	UpdateCertificate(ctx context.Context, id string, certificate UpdateCertificateRequest) (*CertificateResponse, error)
	DeleteCertificate(ctx context.Context, id string) (*CertificateResponse, error)
	GetCertificate(ctx context.Context, id string) (*CertificateResponse, error)
	GetCertificates(ctx context.Context, search *string) ([]*CertificateResponse, error)
}

type CertificateServiceImpl struct {
	CertificateRepository CertificateRepository
	ik                    imagekit.ImageKit
}

func NewService(CertificateRepository CertificateRepository, ik imagekit.ImageKit) CertificateService {
	return &CertificateServiceImpl{CertificateRepository: CertificateRepository, ik: ik}
}

func (p *CertificateServiceImpl) CreateCertificate(ctx context.Context, certificate CreateCertificateRequest) (*CertificateResponse, error) {
	certificateEntity := certificate.ToEntity()

	err := p.CertificateRepository.CreateCertificate(&certificateEntity)
	if err != nil {
		logrus.Errorf("error creating project: %v", err)
		return nil, err
	}

	CertificateResponse := certificateEntity.ToResponse()

	return &CertificateResponse, nil
}

func (p *CertificateServiceImpl) UpdateCertificate(ctx context.Context, id string, updateCertificate UpdateCertificateRequest) (*CertificateResponse, error) {
	if updateCertificate.Title == nil && updateCertificate.ValidUntil == nil && updateCertificate.Url == nil {
		return nil, fmt.Errorf("no input provided")
	}

	certificateEntity := updateCertificate.ToEntity()

	if err := p.CertificateRepository.UpdateCertificate(id, &certificateEntity); err != nil {
		return nil, err
	}

	certificateResponse := certificateEntity.ToResponse()

	return &certificateResponse, nil
}

func (p *CertificateServiceImpl) DeleteCertificate(ctx context.Context, id string) (*CertificateResponse, error) {
	certificateEntity, err := p.CertificateRepository.GetCertificate(id)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Certificate: %+v", certificateEntity)
	err = p.CertificateRepository.DeleteCertificate(id)
	if err != nil {
		return nil, err
	}

	certificateResponse := certificateEntity.ToResponse()

	return &certificateResponse, nil
}

func (p *CertificateServiceImpl) GetCertificate(ctx context.Context, id string) (*CertificateResponse, error) {
	certificateEntity, err := p.CertificateRepository.GetCertificate(id)
	if err != nil {
		return nil, err
	}

	certificateResponse := certificateEntity.ToResponse()

	return &certificateResponse, nil
}

func (p *CertificateServiceImpl) GetCertificates(ctx context.Context, search *string) ([]*CertificateResponse, error) {
	certificateEntities, err := p.CertificateRepository.GetCertificates(search)
	if err != nil {
		return nil, err
	}

	var certificates []*CertificateResponse
	for _, certificateEntity := range certificateEntities {
		projectModel := certificateEntity.ToResponse()
		certificates = append(certificates, &projectModel)
	}

	return certificates, nil
}
