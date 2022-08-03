package cache

import (
	"fmt"
	"main/db/sql"
	dbModels "main/db/sql/models"
	"main/server/api/v1/models"
)

type task struct {
	Id          int32
	Project     int32
	Title       string
	Description string
	Column      int32
	Tags        map[int32]taskAndTags
	Users       map[int32]taskAndUsers
}

type taskAndTags struct {
	Id     int32
	TaskId int32
	TagId  int32
}

type taskAndUsers struct {
	Id     int32
	TaskId int32
	UserId int32
}

func (c internalCache) GetAllTasks() []*models.Task {
	res := []*models.Task{}
	for _, t := range c.Tasks {
		tas := &models.Task{
			Id:          t.Id,
			ProjectId:   t.Project,
			Title:       t.Title,
			Description: t.Description,
			Column:      c.Columns[t.Column].Name,
		}
		for _, user := range t.Users {
			tas.Users = append(tas.Users, user.UserId)
		}
		res = append(res, tas)
	}

	return res
}

func (c internalCache) GetTask(param int32) (*models.Task, error) {
	t, ok := c.Tasks[param]
	if !ok {
		return nil, fmt.Errorf("task with id:%d doesnt exist", param)
	}
	res := &models.Task{
		Id:          t.Id,
		ProjectId:   t.Project,
		Title:       t.Title,
		Description: t.Description,
		Column:      c.Columns[t.Column].Name,
	}
	for _, user := range t.Users {
		res.Users = append(res.Users, user.UserId)
	}

	return res, nil
}

func (c internalCache) GetTasksByProject(params models.TasksByProject) []*models.Task {
	res := []*models.Task{}
	for _, t := range c.Tasks {
		if t.Project == params.Id {
			tas := &models.Task{
				Id:          t.Id,
				ProjectId:   t.Project,
				Title:       t.Title,
				Description: t.Description,
				Column:      c.Columns[t.Column].Name,
			}
			for _, user := range t.Users {
				tas.Users = append(tas.Users, user.UserId)
			}
			res = append(res, tas)
		}
	}

	return res
}

func (c internalCache) UpdateTask(params models.UpdateTask) error {
	t, ok := c.Tasks[params.Id]
	if !ok {
		return fmt.Errorf("task not found")
	}
	t.Project = params.ProjectId
	t.Title = params.Title
	t.Description = params.Description
	for _, col := range c.Columns {
		if params.Column == col.Name {
			t.Column = col.Id
		}
	}
	err := c.updateTasks(t)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) updateTasks(t task) error {
	model := &dbModels.Tasks{
		Id:          t.Id,
		ProjectId:   t.Project,
		Title:       t.Title,
		Description: &t.Description,
		ColumnId:    t.Column,
	}
	err := sql.Update(model)
	if err != nil {
		return err
	}
	c.Tasks[t.Id] = t

	return nil
}

func (c internalCache) DeleteTask(params models.DeleteTask) error {
	t, ok := c.Tasks[params.Id]
	if !ok {
		return fmt.Errorf("task witn id:%d doesnt exist", params.Id)
	}
	err := sql.Delete(&dbModels.Tasks{
		Id: t.Id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) CreateTask(params models.AddTask) error {
	task := &dbModels.Tasks{
		ProjectId:   params.ProjectId,
		Title:       params.Title,
		Description: &params.Description,
	}
	for _, col := range c.Columns {
		if col.Name == params.Column {
			task.ColumnId = col.Id
		}
	}
	if task.ColumnId == 0 {
		return fmt.Errorf("column with name:%s doesnt exist", params.Column)
	}
	err := sql.Upsert(task)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) UpdateTaskPos(params models.UpdateTaskPos) error {
	t, ok := c.Tasks[params.Id]
	if !ok {
		return fmt.Errorf("task not found")
	}
	for _, col := range c.Columns {
		if params.Column == col.Name {
			t.Column = col.Id
		}
	}
	err := c.updateTasks(t)
	if err != nil {
		return err
	}
	return nil
}
