package models

import (
	"github.com/google/uuid"
)

type Column struct {
	Id   uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name string    `json:"name" example:"project name"`
}

type UpdateColumn struct {
	Id   uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name string    `json:"name" form:"name" example:"project name"`
}

type DeleteColumn struct {
	Id uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type AddColumn struct {
	Name string `json:"name" form:"name" example:"column name"`
}
