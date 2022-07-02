package models

import (
	"time"
)

type Project struct {
	Id           int32     `json:"id" format:"integer"`
	ParentId     int32     `json:"parentId" format:"integer"`
	Name         string    `json:"name" example:"project name"`
	Description  string    `json:"description" example:"description"`
	CreationDate time.Time `json:"creationDate" format:"datetime"`
	OwnerId      int32     `json:"ownerId" format:"integer"`
}

type UpdateProject struct {
	Id          int32  `json:"id" form:"id" format:"integer"`
	ParentId    int32  `json:"parentId" form:"parentId" format:"integer"`
	Name        string `json:"name" form:"name" example:"project name"`
	Description string `json:"description" form:"description" example:"description"`
	OwnerId     int32  `json:"ownerId" form:"ownerId" format:"integer"`
}

type DeleteProject struct {
	Id int32 `json:"id" form:"id" format:"integer"`
}

type GetProject struct {
	Id int32 `json:"id" form:"id" format:"integer"`
}

type AddProject struct {
	ParentId    int32  `json:"parentId" form:"parentId" format:"integer"`
	Name        string `json:"name" form:"name" example:"project name"`
	Description string `json:"description" form:"description" example:"description"`
	OwnerId     int32  `json:"ownerId" form:"ownerId" format:"integer"`
}

type TreeNode struct {
	Title    string     `json:"title" example:"Project name"`
	Key      int32      `json:"key"`
	Children []TreeNode `json:"children"`
}
