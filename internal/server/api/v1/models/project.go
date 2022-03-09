package models

import (
	"github.com/golang-sql/civil"
	"github.com/google/uuid"
)

type Project struct {
	Id           uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ParentId     uuid.UUID `json:"parentId" example:"550e8400-e29b-41d4-a716-44665544000" format:"uuid"`
	Name         string    `json:"name" example:"project name"`
	Description  string    `json:"description" example:"example"`
	CreationDate civil.DateTime
}
