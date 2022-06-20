package sql

import (
	"main/database/sql/models"
	apiModels "main/server/api/v1/models"
)

func (DbProvider DatabaseProvider) CreateNewColumn(params apiModels.AddColumn) error {
	s := &models.Columns{
		Name: params.Name,
	}

	err := DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}

func (DbProvider DatabaseProvider) GetAllColumns() ([]apiModels.Column, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.ColumnsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.Column
	for _, s := range from {

		s := s.(*models.Columns)
		q := apiModels.Column{
			Id:   s.Id,
			Name: s.Name,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) DeleteColumn(params apiModels.DeleteColumn) error {

	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.ColumnsTable, params.Id)
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

func (DbProvider DatabaseProvider) UpdateColumn(params apiModels.UpdateColumn) error {
	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.ColumnsTable, params.Id)
	if err != nil {
		return err //TODO
	}
	rec := from.(*models.Columns)

	s := &models.Columns{
		Name: NilCheck(params.Name, rec.Name).(string),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}
