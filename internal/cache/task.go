package cache

import (
	"main/db/sql"
	dbModels "main/db/sql/models"
	"main/server/api/v1/models"
)

func (c cache) GetAllTasks() []models.Task {
	res := []models.Task{}
	for _, t := range c.Tasks {
		tas := models.Task{
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

func (c cache) GetTask(param int32) models.Task {
	t := c.Tasks[param]
	res := models.Task{
		Id:          t.Id,
		ProjectId:   t.Project,
		Title:       t.Title,
		Description: t.Description,
		Column:      c.Columns[t.Column].Name,
	}
	for _, user := range t.Users {
		res.Users = append(res.Users, user.UserId)
	}

	return res
}

func (c cache) GetTasksByProject(params models.TasksByProject) []models.Task {
	res := []models.Task{}
	for _, t := range c.Tasks {
		if t.Project == params.Id {
			tas := models.Task{
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

func (c cache) UpdateTask(params models.UpdateTask) error {
	upd := map[int32]task{}
	upd[params.Id] = task{
		Id:          params.Id,
		Project:     params.ProjectId,
		Title:       params.Title,
		Description: params.Description,
	}

	err := c.updateTasks(upd)
	if err != nil {
		return err
	}
	return nil
}

func (c cache) updateTasks(upd map[int32]task) error {
	for id, task := range upd {
		model := &dbModels.Tasks{
			Id:          task.Id,
			ProjectId:   task.Project,
			Title:       task.Title,
			Description: &task.Description,
			ColumnId:    task.Column,
		}
		err := sql.Update(model)
		if err != nil {
			return err
		}
		c.Tasks[id] = task
	}
	return nil
}
