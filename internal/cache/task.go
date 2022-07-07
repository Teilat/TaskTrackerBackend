package cache

import (
	"main/db/sql"
	dbModels "main/db/sql/models"
	"main/server/api/v1/models"
	"reflect"
)

func (c cache) GetAllTasks() []models.Task {
	res := []models.Task{}

	for _, t := range c.Tasks {
		res = append(res, models.Task{
			Id:          t.Id,
			ProjectId:   t.Project,
			Title:       t.Title,
			Description: t.Description,
			Column:      c.Columns[t.Column].Name,
			Users:       t.Users,
		})
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
		Users:       t.Users,
	}

	return res
}

func (c cache) GetTasksByProject(params models.TasksByProject) []models.Task {
	res := []models.Task{}

	for _, t := range c.Tasks {
		if t.Project == params.Id {
			res = append(res, models.Task{
				Id:          t.Id,
				ProjectId:   t.Project,
				Title:       t.Title,
				Description: t.Description,
				Column:      c.Columns[t.Column].Name,
				Users:       t.Users,
			})
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
		Tags:        params.Tags,
		Users:       params.Users,
	}

	err := c.updateTasks(upd)
	if err != nil {
		return err
	}
	return nil
}

func (c cache) updateTasks(upd map[int32]task) error {
	for id, task := range upd {
		var addedTags, removedTags []int32
		model := &dbModels.Tasks{
			Id:          task.Id,
			ProjectId:   task.Project,
			Title:       task.Title,
			Description: &task.Description,
			ColumnId:    task.Column,
		}
		if !reflect.DeepEqual(c.Tasks[id].Tags, task.Tags) {
			addedTags = difference(c.Tasks[id].Tags, task.Tags)
			removedTags = difference(task.Tags, c.Tasks[id].Tags)
		}
		err := sql.Update(model)
		if err != nil {
			return err
		}
		if len(addedTags) > 0 {
			for _, tagId := range addedTags {
				err = sql.Upsert(&dbModels.TaskAndTags{
					TaskId: model.Id,
					TagId:  tagId,
				})
				if err != nil {
					return err
				}
			}
		}
		if len(removedTags) > 0 {
			for _, tagId := range removedTags {
				err = sql.Delete(&dbModels.TaskAndTags{
					TaskId: model.Id,
					TagId:  tagId,
				})
				if err != nil {
					return err
				}
			}
		}
		c.Tasks[id] = task
	}
	return nil
}

// a:старое, b:новое
func difference(a, b []int32) []int32 {
	mb := make(map[int32]interface{}, 0)
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int32
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
