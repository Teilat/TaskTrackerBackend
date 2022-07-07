package models

type Task struct {
	Id          int32   `json:"id" format:"integer"`
	ProjectId   int32   `json:"projectId" format:"integer"`
	Title       string  `json:"title" example:"task name"`
	Description string  `json:"description" example:"description"`
	Column      string  `json:"column" form:"column" format:"string"`
	Users       []int32 `json:"users" form:"users" format:"[]integer"`
	Tags        []int32 `json:"tags" form:"tags" format:"[]integer"`
}

type AddTask struct {
	ProjectId   int32  `json:"projectId" form:"projectId" format:"integer"`
	Title       string `json:"title" form:"title" example:"task name"`
	Description string `json:"description" form:"description" example:"description"`
	Column      string `json:"column" form:"column" format:"string"`
}

type UpdateTask struct {
	Id          int32   `json:"id" form:"id" format:"integer"`
	ProjectId   int32   `json:"projectId" form:"projectId" format:"integer"`
	Title       string  `json:"title" form:"title" example:"task name"`
	Description string  `json:"description" form:"description" example:"description"`
	Column      string  `json:"column" form:"column" format:"string"`
	Users       []int32 `json:"users" form:"users" format:"[]integer"`
	Tags        []int32 `json:"tags" form:"tags" format:"[]integer"`
}

type DeleteTask struct {
	Id int32 `form:"id" format:"integer"`
}

type UpdateTaskPos struct {
	Id     int32  `form:"id" format:"integer"`
	Column string `json:"column" form:"column" format:"string"`
}

type TasksByProject struct {
	Id int32 `form:"id" format:"integer"`
}
