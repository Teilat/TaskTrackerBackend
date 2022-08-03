package cache

import (
	"fmt"
	"main/db/sql"
	dbModels "main/db/sql/models"
	"main/server/api/v1/models"
)

type column struct {
	Id   int32
	Name string
}

func (c internalCache) CreateColumn(params models.AddColumn) error {
	col := &dbModels.Columns{
		Name: params.Name,
	}
	err := sql.Upsert(col)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) GetAllColumns() *[]models.Column {
	var res []models.Column
	for _, col := range c.Columns {
		res = append(res, models.Column{
			Id:   col.Id,
			Name: col.Name,
		})
	}
	return &res
}

func (c internalCache) UpdateColumn(params models.UpdateColumn) error {
	col := c.Columns[params.Id]
	col.Name = params.Name
	err := c.updateColumn(col)
	if err != nil {
		return err
	}
	return nil
}

func (c internalCache) updateColumn(col column) error {
	model := &dbModels.Columns{
		Id:   col.Id,
		Name: col.Name,
	}
	err := sql.Update(model)
	if err != nil {
		return err
	}
	c.Columns[col.Id] = col
	return nil
}

func (c internalCache) DeleteColumn(params models.DeleteColumn) error {
	col, ok := c.Columns[params.Id]
	if ok {
		return fmt.Errorf("column not found")
	}
	model := &dbModels.Columns{
		Id: col.Id,
	}
	err := sql.Delete(model)
	if err != nil {
		return err
	}
	delete(c.Columns, col.Id)
	return nil
}
