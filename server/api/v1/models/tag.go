package models

type Tag struct {
	Id    int32  `json:"id" format:"integer"`
	Name  string `json:"name" example:"tag name"`
	Color string `json:"color" example:"FFFFFF"`
}

type AddTag struct {
	Name  string `form:"name" json:"name" example:"tag name"`
	Color string `form:"color" json:"color" example:"FFFFFF"`
}

type UpdateTag struct {
	Id    int32  `form:"id" json:"id" format:"integer"`
	Name  string `form:"name" json:"name" example:"tag name"`
	Color string `form:"color" json:"color" example:"FFFFFF"`
}

type DeleteTag struct {
	Id int32 `form:"id" json:"id" format:"integer"`
}

type TagsByTask struct {
	Id int32 `json:"id" form:"id" format:"integer"`
}

type TagsByProject struct {
	Id int32 `json:"id" form:"id" format:"integer"`
}
