package sql

import (
	"main/db/sql/models"
	apiModels "main/server/api/v1/models"
)

func (db *reform.DB) CreateNewTag(params apiModels.AddTag) error {
	s := &models.Tags{
		Name:  params.Name,
		Color: params.Color,
	}

	err := db.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}

func (DbProvider DatabaseProvider) GetAllTags() ([]apiModels.Tag, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.TagsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apiModels.Tag
	for _, s := range from {

		s := s.(*models.Tags)
		q := apiModels.Tag{
			Id:    s.Id,
			Name:  s.Name,
			Color: s.Color,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) DeleteTag(params apiModels.DeleteTag) error {

	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TagsTable, params.Id)
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

func (DbProvider DatabaseProvider) UpdateTag(params apiModels.UpdateTag) error {
	from, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TagsTable, params.Id)
	if err != nil {
		return err //TODO
	}
	rec := from.(*models.Tags)

	s := &models.Tags{
		Name:  NilCheck(params.Name, rec.Name).(string),
		Color: NilCheck(params.Color, rec.Color).(string),
	}

	err = DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}

func (DbProvider DatabaseProvider) GetTagsByTask(params apiModels.TagsByTask) ([]apiModels.Tag, error) {

	tagsInTask, err := DbProvider.DB.SelectAllFrom(models.TaskAndTagsTable, "TaskId", params.Id)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}

	from, err := DbProvider.DB.SelectAllFrom(models.TagsTable, "Id", tagsInTask)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}

	var list []apiModels.Tag
	for _, s := range from {

		s := s.(*models.Tags)
		q := apiModels.Tag{
			Id:    s.Id,
			Name:  s.Name,
			Color: s.Color,
		}
		list = append(list, q)
	}
	return list, nil

}
