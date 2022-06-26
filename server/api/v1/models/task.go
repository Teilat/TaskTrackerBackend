package models

type Task struct {
	Id          int32   `json:"id" example:"15" format:"integer"`
	ProjectId   int32   `json:"projectId" example:"2" format:"integer"`
	Title       string  `json:"title" example:"task name"`
	Description string  `json:"description" example:"example task"`
	Column      string  `json:"column" form:"column" example:"To Do" format:"string"`
	Users       []int32 `json:"users" form:"users" example:"12,24,36,1" format:"[]integer"`
}

type AddTask struct {
	ProjectId   int32  `json:"projectId" form:"projectId" example:"2" format:"integer"`
	Title       string `json:"Title" form:"Title" example:"task name"`
	Description string `json:"Description" form:"Description" example:"example task"`
	Column      string `json:"column" form:"column" example:"To Do" format:"string"`
}

type UpdateTask struct {
	Id          int32   `json:"id" form:"id" example:"15" format:"integer"`
	ProjectId   int32   `json:"projectId" form:"projectId" example:"2" format:"integer"`
	Title       string  `json:"Title" form:"Title" example:"task name"`
	Description string  `json:"Description" form:"Description" example:"example task"`
	Column      string  `json:"column" form:"column" example:"To Do" format:"string"`
	Users       []int32 `json:"users" form:"users" example:"12,24,36,1" format:"[]integer"`
}

type DeleteTask struct {
	Id int32 `form:"id" example:"15" format:"integer"`
}

type UpdateTaskPos struct {
	Id     int32  `form:"id" example:"15" format:"integer"`
	Column string `json:"column" form:"column" example:"To Do" format:"string"`
}

type TasksByProject struct {
	Id int32 `form:"id" example:"15" format:"integer"`
}
