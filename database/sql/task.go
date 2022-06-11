package sql

import (
	"github.com/google/uuid"
	"main/database/sql/models"
	apiModels "main/server/api/v1/models"
	"strings"
)

func (DbProvider DatabaseProvider) CreateNewTask(params apiModels.AddTask) error {
	newUuid := uuid.New()

	columns, err := DbProvider.DB.SelectAllFrom(models.ColumnsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	var c uuid.UUID

	for _, column := range columns {
		column := column.(*models.Columns)
		if params.Column == column.Name {
			c = column.ID
		}
	}

	s := &models.Tasks{
		Id:          newUuid,
		ProjectId:   params.ProjectId,
		Title:       params.Title,
		Description: &params.Description,
		ColumnId:    c,
	}

	err = DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Fatalln(err)
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

	columns, err := DbProvider.DB.SelectAllFrom(models.ColumnsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}

	var list []apiModels.Task
	for _, s := range from {
		s := s.(*models.Tasks)

		var c string
		for _, column := range columns {
			column := column.(*models.Columns)
			if s.ColumnId == column.ID {
				c = column.Name
			}
		}
		q := apiModels.Task{
			Id:          s.Id,
			ProjectId:   s.ProjectId,
			Title:       s.Title,
			Description: *s.Description,
			Column:      c,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) GetAllTasksByProject(params apiModels.TasksByProject) ([]apiModels.Task, error) {
	from, err := DbProvider.DB.SelectAllFrom(models.TasksTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}

	columns, err := DbProvider.DB.SelectAllFrom(models.ColumnsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}

	list := []apiModels.Task{}
	id := strings.ToLower(params.Id)

	for _, s := range from {
		s := s.(*models.Tasks)
		if s.ProjectId.String() != id {
			continue
		}
		var columnsName string
		for _, column := range columns {
			column := column.(*models.Columns)
			if s.ColumnId == column.ID {
				columnsName = column.Name
			}
		}

		q := apiModels.Task{
			Id:          s.Id,
			ProjectId:   s.ProjectId,
			Title:       s.Title,
			Description: *s.Description,
			Column:      columnsName,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) DeleteTask(params apiModels.DeleteTask) error {

	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TasksTable, params.Id)
	if err != nil {
		DbProvider.DbLogger.Fatal(err)
		return err
	}

	err = DbProvider.DB.Delete(record)
	if err != nil {
		DbProvider.DbLogger.Fatal(err)
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
		Id:          NilCheck(params.Id, rec.Id).(uuid.UUID),
		ProjectId:   NilCheck(params.ProjectId, rec.ProjectId).(uuid.UUID),
		Title:       NilCheck(params.Title, rec.Title).(string),
		Description: NilCheck(params.Description, rec.Description).(*string),
		ColumnId:    rec.ColumnId,
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Fatal(err)
		return err
	}

	return nil
}
