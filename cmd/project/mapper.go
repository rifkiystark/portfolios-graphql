package project

func (cp *CreateProjectRequest) ToEntity() Project {
	return Project{
		Title:          cp.Title,
		AdditionalInfo: cp.AdditionalInfo,
		Description:    cp.Description,
	}
}

func (up *UpdateProjectRequest) ToEntity() Project {
	projectEntity := Project{}

	if up.Title != nil {
		projectEntity.Title = *up.Title
	}

	if up.Description != nil {
		projectEntity.Description = *up.Description
	}

	if up.AdditionalInfo != nil {
		projectEntity.AdditionalInfo = up.AdditionalInfo
	}

	return projectEntity
}

func (cp *Project) ToResponse() ProjectResponse {
	return ProjectResponse{
		ID:             cp.ID.Hex(),
		Title:          cp.Title,
		AdditionalInfo: cp.AdditionalInfo,
		Description:    cp.Description,
		ImageURL:       cp.ImageURL,
	}
}
