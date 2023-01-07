package ipr

func (cp *CreateIPRRequest) ToEntity() IPR {
	return IPR{
		Title:       cp.Title,
		PublishedAt: cp.PublishedAt,
		Description: cp.Description,
		Url:         cp.Url,
	}
}

func (up *UpdateIPRRequest) ToEntity() IPR {
	iprEntity := IPR{}

	if up.Title != nil {
		iprEntity.Title = *up.Title
	}

	if up.Description != nil {
		iprEntity.Description = *up.Description
	}

	if up.PublishedAt != nil {
		iprEntity.PublishedAt = *up.PublishedAt
	}

	if up.Url != nil {
		iprEntity.Url = *up.Url
	}

	return iprEntity
}

func (cp *IPR) ToResponse() IPRResponse {
	return IPRResponse{
		ID:          cp.ID.Hex(),
		Title:       cp.Title,
		PublishedAt: cp.PublishedAt,
		Description: cp.Description,
		Url:         cp.Url,
	}
}
