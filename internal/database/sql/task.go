package sql

import (
	"github.com/google/uuid"
	"main/internal/database/sql/models"
	apiModels "main/internal/server/api/v1/models"
)

func (DbProvider DatabaseProvider) CreateNewTask(params apiModels.AddTask) error {
	newUuid := uuid.New()
	s := &models.Tasks{
		Id:              newUuid,
		ProjectId:       params.ProjectId,
		TaskTitle:       params.TaskTitle,
		TaskDescription: params.TaskDescription,
	}

	err := DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}

func (DbProvider DatabaseProvider) GetAllTasks() ([]apiModels.Task, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.TasksTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.Task
	for _, s := range from {

		s := s.(*models.Tasks)
		q := apiModels.Task{
			Id:              s.Id,
			ProjectId:       s.ProjectId,
			TaskTitle:       s.TaskTitle,
			TaskDescription: s.TaskDescription,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) DeleteTask(params apiModels.DeleteTask) error {

	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TasksTable, params.Id)
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

func (DbProvider DatabaseProvider) UpdateTask(params apiModels.UpdateTask) error {

	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TasksTable, params.Id)
	if err != nil {
		return err //TODO
	}
	rec := from.(*models.Tasks)

	s := &models.Tasks{
		Id:              NilCheck(params.Id, rec.Id).(uuid.UUID),
		ProjectId:       NilCheck(params.ProjectId, rec.ProjectId).(uuid.UUID),
		TaskTitle:       NilCheck(params.TaskTitle, rec.TaskTitle).(string),
		TaskDescription: NilCheck(params.TaskDescription, rec.TaskDescription).(string),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}
