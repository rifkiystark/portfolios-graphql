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

func (cp *Project) FillModelByDBEntity(dp database.Project) {
	cp.Title = dp.Title
	cp.AdditionalInfo = dp.AdditionalInfo
	cp.Description = dp.Description
	cp.ImageURL = dp.ImageURL
}
