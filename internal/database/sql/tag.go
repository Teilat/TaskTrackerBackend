package sql

import (
	"github.com/google/uuid"
	apimodels "main/internal/api/v1/models"
	"main/internal/database/sql/models"
)

func (DbProvider DatabaseProvider) CreateNewTag(tagName, tagColor string) error {
	newUuid := uuid.New()
	s := &models.Tags{
		Id:       newUuid,
		TagName:  tagName,
		TagColor: tagColor,
	}

	err := DbProvider.DB.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}
func (DbProvider DatabaseProvider) GetAllTags() ([]apimodels.Tag, error) {

	from, err := DbProvider.DB.SelectAllFrom(models.TagsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}
	var list []apimodels.Tag
	for _, s := range from {
		q := apimodels.Tag{
			Id:    s.(*models.Tags).Id,
			Name:  s.(*models.Tags).TagName,
			Color: s.(*models.Tags).TagColor,
		}
		list = append(list, q)
	}
	return list, nil
}

func (DbProvider DatabaseProvider) RemoveTag(PKey uuid.UUID) error {

	record, err := DbProvider.DB.FindByPrimaryKeyFrom(models.TagsTable, PKey)
	if err != nil {
		return err
	}

	err = DbProvider.DB.Delete(record)
	if err != nil {
		return err
	}

	return nil
}
