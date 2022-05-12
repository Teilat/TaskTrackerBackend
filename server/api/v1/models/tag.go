package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	Id    uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name  string    `json:"name" example:"tag name"`
	Color string    `json:"color" example:"FFFFFF"`
}

type AddTag struct {
	Name  string `form:"name" json:"name" example:"tag name"`
	Color string `form:"color" json:"color" example:"FFFFFF"`
}

type UpdateTag struct {
	Id    uuid.UUID `form:"id" json:"id" example:"f40312cd-5995-ec11-b909-0242ac120002" format:"uuid"`
	Name  string    `form:"name" json:"name" example:"tag name"`
	Color string    `form:"color" json:"color" example:"FFFFFF"`
}

type DeleteTag struct {
	Id uuid.UUID `form:"id" json:"id" example:"f40312cd-5995-ec11-b909-0242ac120002" format:"uuid"`
}

type TagsByTask struct {
	TaskId uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type TagsByProject struct {
	TaskId uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}
