package models

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	Id           uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ParentId     uuid.UUID `json:"parentId" example:"550e8400-e29b-41d4-a716-44665544000" format:"uuid"`
	Name         string    `json:"name" example:"project name"`
	Description  string    `json:"description" example:"example"`
	CreationDate time.Time `json:"creationDate" format:"datetime"`
	OwnerId      uuid.UUID `json:"ownerId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type UpdateProject struct {
	Id          uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ParentId    uuid.UUID `json:"parentId" form:"parentId" example:"550e8400-e29b-41d4-a716-44665544000" format:"uuid"`
	Name        string    `json:"name" form:"name" example:"project name"`
	Description *string   `json:"description" form:"description" example:"example"`
	OwnerId     uuid.UUID `json:"ownerId" form:"ownerId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type DeleteProject struct {
	Id uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type AddProject struct {
	ParentId    uuid.UUID `json:"parentId" form:"parentId" example:"550e8400-e29b-41d4-a716-44665544000" format:"uuid"`
	Name        string    `json:"name" form:"name" example:"project name"`
	Description string    `json:"description" form:"description" example:"example"`
	OwnerId     uuid.UUID `json:"ownerId" form:"ownerId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}
