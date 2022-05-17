package sql

import (
	"github.com/google/uuid"
	"main/database/sql/models"
	apiModels "main/server/api/v1/models"
	"time"
)

func (DbProvider DatabaseProvider) CreateProject(params apiModels.AddProject) error {
	newUuid := uuid.New()
	var s *models.Projects
	if params.ParentId == uuid.Nil {
		s = &models.Projects{
			Id:                 newUuid,
			ParentId:           nil,
			ProjectName:        params.Name,
			ProjectDescription: params.Description,
			CreationDate:       time.Now(),
			OwnerId:            params.OwnerId,
		}
	} else {
		s = &models.Projects{
			Id:                 newUuid,
			ParentId:           &params.ParentId,
			ProjectName:        params.Name,
			ProjectDescription: params.Description,
			CreationDate:       time.Now(),
			OwnerId:            params.OwnerId,
		}
	}

	err := DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}

func (DbProvider DatabaseProvider) GetAllProjects() ([]apiModels.Project, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.ProjectsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.Project
	for _, s := range from {

		s := s.(*models.Projects)
		q := apiModels.Project{
			Id:           s.Id,
			Name:         s.ProjectName,
			Description:  s.ProjectDescription,
			CreationDate: s.CreationDate,
			OwnerId:      s.OwnerId,
		}
		if s.ParentId != nil {
			q.ParentId = *s.ParentId
		} else {
			q.ParentId = uuid.Nil
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) DeleteProject(params apiModels.DeleteProject) error {
	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.ProjectsTable, params.Id)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	err = DbProvider.DB.Delete(record)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}

func (DbProvider DatabaseProvider) UpdateProject(params apiModels.UpdateProject) error {

	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.ProjectsTable, params.Id)
	if err != nil {
		return err //TODO
	}
	rec := from.(*models.Projects)

	s := &models.Projects{
		Id:                 NilCheck(params.Id, rec.Id).(uuid.UUID),
		ParentId:           NilCheck(&params.ParentId, &rec.ParentId).(*uuid.UUID),
		ProjectName:        NilCheck(params.Name, rec.ProjectName).(string),
		ProjectDescription: NilCheck(params.Description, rec.ProjectDescription).(string),
		OwnerId:            NilCheck(params.OwnerId, rec.OwnerId).(uuid.UUID),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}
