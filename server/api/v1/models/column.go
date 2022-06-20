package models

type Column struct {
	Id   int32  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"integer"`
	Name string `json:"name" example:"project name"`
}

type UpdateColumn struct {
	Id   int32  `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"integer"`
	Name string `json:"name" form:"name" example:"project name"`
}

type DeleteColumn struct {
	Id int32 `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"integer"`
}

type AddColumn struct {
	Name string `json:"name" form:"name" example:"column name"`
}
