package models

import (
	"github.com/google/uuid"
)

type Tag struct {
	Id    uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name  string    `json:"name" example:"tag name"`
	Color string    `json:"color" example:"ffffff"`
}

type AddTag struct {
	Name  string `json:"name" example:"tag name"`
	Color string `json:"color" example:"ffffff"`
}

type UpdateTag struct {
	Name  string `json:"name" example:"tag name"`
	Color string `json:"color" example:"ffffff"`
}
