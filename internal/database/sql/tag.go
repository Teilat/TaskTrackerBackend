package sql

import (
	"github.com/google/uuid"
	"gopkg.in/reform.v1"
	"main/internal/database/sql/models"
)

func (DbProvider DatabaseProvider) CreateNewTag(tagName, tagColor string) error {
	newUuid := uuid.New()
	s := &models.Tags{
		Id:       newUuid,
		TagName:  tagName,
		TagColor: tagColor,
	}

	err := DbProvider.Db.Save(s)
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return err
	}
	return nil
}
func (DbProvider DatabaseProvider) GetAllTags() ([]reform.Struct, error) {

	from, err := DbProvider.Db.SelectAllFrom(models.TagsTable, "")
	if err != nil {
		DbProvider.DbLogger.Println(err)
		return nil, err
	}

	return from, nil
}

func (DbProvider DatabaseProvider) RemoveTag(PKey uuid.UUID) error {

	record, err := DbProvider.Db.FindByPrimaryKeyFrom(models.TagsTable, PKey)
	if err != nil {
		return err
	}

	err = DbProvider.Db.Delete(record)
	if err != nil {
		return err
	}

	return nil
}
