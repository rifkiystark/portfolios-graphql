package model

import (
	"github.com/rifkiystark/portfolios-api/database"
)

func (cp *CreateProject) ToProjectEntity() database.Project {
	return database.Project{
		Title:          cp.Title,
		AdditionalInfo: cp.AdditionalInfo,
		Description:    cp.Description,
	}
}

func (p *Project) FillModelByDBEntity(dp database.Project) {
	p.Title = dp.Title
	p.AdditionalInfo = dp.AdditionalInfo
	p.Description = dp.Description
	p.ImageURL = dp.ImageURL
}
