package sql

import (
	"github.com/google/uuid"
	"main/internal/database/sql/models"
	apiModels "main/internal/server/api/v1/models"
)

func (DbProvider DatabaseProvider) CreateNewTag(params apiModels.AddTag) error {
	newUuid := uuid.New()
	s := &models.Tags{
		Id:       newUuid,
		TagName:  params.Name,
		TagColor: params.Color,
	}

	err := DbProvider.DB.Save(s)
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
			Name:  s.TagName,
			Color: s.TagColor,
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

	s := &models.Tags{
		Id:       params.Id,
		TagName:  params.Name,
		TagColor: params.Color,
	}

	err := DbProvider.DB.Update(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}

	return nil
}

func (DbProvider DatabaseProvider) GetTagsByTask(params apiModels.TagsByTask) ([]apiModels.Tag, error) {

	tagsInTask, err := DbProvider.DB.SelectAllFrom(models.TaskAndTagsTable, "TaskId", params.TaskId)
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
			Name:  s.TagName,
			Color: s.TagColor,
		}
		list = append(list, q)
	}
	return list, nil

}
