package sql

import (
	"main/db/sql/models"
	apiModels "main/server/api/v1/models"
)

func (DbProvider DatabaseProvider) CreateNewTask(params apiModels.AddTask) error {

	columns, err := DbProvider.DB.SelectAllFrom(models.ColumnsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	var c int32

	for _, column := range columns {
		column := column.(*models.Columns)
		if params.Column == column.Name {
			c = column.Id
		}
	}

	s := &models.Tasks{
		ProjectId:   params.ProjectId,
		Title:       params.Title,
		Description: &params.Description,
		ColumnId:    c,
	}

	//err = Upsert[*models.Tasks]([]*models.Tasks{s})
	//if err != nil {
	//	fmt.Println(err)
	//}

	err = DbProvider.DB.Save(s)
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
			if s.ColumnId == column.Id {
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

	for _, s := range from {
		s := s.(*models.Tasks)
		if s.ProjectId != params.Id {
			continue
		}
		var columnsName string
		for _, column := range columns {
			column := column.(*models.Columns)
			if s.ColumnId == column.Id {
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

func (DbProvider DatabaseProvider) UpdateTaskPos(params apiModels.UpdateTaskPos) error {
	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TasksTable, params.Id)
	if err != nil {
		DbProvider.DbLogger.Fatal(err)
		return err
	}
	rec := record.(*models.Tasks)

	columns, err := DbProvider.DB.SelectAllFrom(models.ColumnsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	for _, column := range columns {
		column := column.(*models.Columns)
		if params.Column == column.Name {
			rec.ColumnId = column.Id
		}
	}

	err = DbProvider.DB.Update(rec)
	if err != nil {
		DbProvider.DbLogger.Fatal(err)
		return err
	}

	return nil
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
		ProjectId:   NilCheck(params.ProjectId, rec.ProjectId).(int32),
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
