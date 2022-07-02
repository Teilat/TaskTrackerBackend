package models

type Column struct {
	Id   int32  `json:"id" format:"integer"`
	Name string `json:"name" example:"To Do"`
}

type UpdateColumn struct {
	Id   int32  `json:"id" form:"id" format:"integer"`
	Name string `json:"name" form:"name" example:"To Do"`
}

type DeleteColumn struct {
	Id int32 `json:"id" form:"id" format:"integer"`
}

type AddColumn struct {
	Name string `json:"name" form:"name" example:"To Do"`
}
