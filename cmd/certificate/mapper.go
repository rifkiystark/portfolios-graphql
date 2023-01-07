package certificate

func (cc *CreateCertificateRequest) ToEntity() Certificate {
	return Certificate{
		Title:      cc.Title,
		ValidUntil: cc.ValidUntil,
		Url:        cc.Url,
	}
}

func (uc *UpdateCertificateRequest) ToEntity() Certificate {
	certificateEntity := Certificate{}

	if uc.Title != nil {
		certificateEntity.Title = *uc.Title
	}

	if uc.ValidUntil != nil {
		certificateEntity.ValidUntil = *uc.ValidUntil
	}

	if uc.Url != nil {
		certificateEntity.Url = *uc.Url
	}

	return certificateEntity
}

func (c *Certificate) ToResponse() CertificateResponse {
	return CertificateResponse{
		ID:         c.ID.Hex(),
		Title:      c.Title,
		ValidUntil: c.ValidUntil,
		Url:        c.Url,
	}
}
