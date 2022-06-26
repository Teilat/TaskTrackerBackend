package models

import (
	"time"
)

type Project struct {
	Id           int32     `json:"id" example:"2" format:"integer"`
	ParentId     int32     `json:"parentId" example:"1" format:"integer"`
	Name         string    `json:"name" example:"project name"`
	Description  string    `json:"description" example:"example"`
	CreationDate time.Time `json:"creationDate" format:"datetime"`
	OwnerId      int32     `json:"ownerId" example:"1" format:"integer"`
}

type UpdateProject struct {
	Id          int32   `json:"id" form:"id" example:"2" format:"integer"`
	ParentId    int32   `json:"parentId" form:"parentId" example:"1" format:"integer"`
	Name        string  `json:"name" form:"name" example:"project name"`
	Description *string `json:"description" form:"description" example:"example"`
	OwnerId     int32   `json:"ownerId" form:"ownerId" example:"1" format:"integer"`
}

type DeleteProject struct {
	Id int32 `json:"id" form:"id" example:"2" format:"integer"`
}

type GetProject struct {
	Id int32 `json:"id" form:"id" example:"2" format:"integer"`
}

type AddProject struct {
	ParentId    int32  `json:"parentId" form:"parentId" example:"2" format:"integer"`
	Name        string `json:"name" form:"name" example:"project name"`
	Description string `json:"description" form:"description" example:"example"`
	OwnerId     int32  `json:"ownerId" form:"ownerId" example:"1" format:"integer"`
}

type TreeNode struct {
	Title    string     `json:"title" example:"ProjectName"`
	Key      int32      `json:"key"`
	Children []TreeNode `json:"children"`
}
