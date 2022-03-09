package sql

import (
	"main/internal/database/sql/models"
	apiModels "main/internal/server/api/v1/models"
)

func (DbProvider DatabaseProvider) GetAllProjects() ([]apiModels.Tag, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.ProjectsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.Tag
	for _, s := range from {

		q := apiModels.Tag{
			Id:    s.(*models.Tags).Id,
			Name:  s.(*models.Tags).TagName,
			Color: s.(*models.Tags).TagColor,
		}
		list = append(list, q)
	}
	return list, nil
}
