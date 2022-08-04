package cache

import (
	"fmt"
	"main/db/sql"
	dbModels "main/db/sql/models"
	"main/server/api/v1/models"
)

type tag struct {
	Id    int32
	Name  string
	Color string
}

func (c internalCache) CreateTag(params models.AddTag) error {
	t := &dbModels.Tags{
		Name:  params.Name,
		Color: params.Color,
	}
	err := sql.Upsert(t)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) GetAllTags() *[]models.Tag {
	var res []models.Tag
	for _, col := range c.Tags {
		res = append(res, models.Tag{
			Id:    col.Id,
			Name:  col.Name,
			Color: col.Color,
		})
	}
	return &res
}

func (c internalCache) UpdateTag(params models.UpdateTag) error {
	t := c.Tags[params.Id]
	t.Name = params.Name
	t.Color = params.Color
	err := c.updateTag(t)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) updateTag(t tag) error {
	model := &dbModels.Tags{
		Id:    t.Id,
		Name:  t.Name,
		Color: t.Color,
	}
	err := sql.Update(model)
	if err != nil {
		return err
	}
	c.Tags[t.Id] = t
	return nil
}

func (c internalCache) DeleteTag(params models.DeleteTag) error {
	col, ok := c.Tags[params.Id]
	if ok {
		return fmt.Errorf("tag not found")
	}
	model := &dbModels.Tags{
		Id: col.Id,
	}
	err := sql.Delete(model)
	if err != nil {
		return err
	}
	delete(c.Tags, col.Id)
	return nil
}

func (c internalCache) GetTagsByTask(params models.TagsByTask) *[]models.Tag {
	var res []models.Tag
	for _, t := range c.Tasks[params.Id].Tags {
		tags := c.Tags[t.TagId]
		res = append(res, models.Tag{
			Id:    tags.Id,
			Name:  tags.Name,
			Color: tags.Color,
		})
	}
	return &res
}
