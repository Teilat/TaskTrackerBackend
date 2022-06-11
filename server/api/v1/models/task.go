package models

import "github.com/google/uuid"

type Task struct {
	Id          uuid.UUID   `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ProjectId   uuid.UUID   `json:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Title       string      `json:"title" example:"task name"`
	Description string      `json:"description" example:"example task"`
	Column      string      `json:"column" form:"column" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Users       []uuid.UUID `json:"users" form:"users" example:"[550e8400-e29b-41d4-a716-446655440000,550e8400-e29b-41d4-a716-446655440000]" format:"[]uuid"`
}

type AddTask struct {
	ProjectId   uuid.UUID `json:"projectId" form:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Title       string    `json:"taskTitle" form:"taskTitle" example:"task name"`
	Description string    `json:"taskDescription" form:"taskDescription" example:"example task"`
	Column      string    `json:"column" form:"column" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type UpdateTask struct {
	Id          uuid.UUID   `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ProjectId   uuid.UUID   `json:"projectId" form:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Title       string      `json:"taskTitle" form:"taskTitle" example:"task name"`
	Description string      `json:"taskDescription" form:"taskDescription" example:"example task"`
	Column      string      `json:"column" form:"column" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Users       []uuid.UUID `json:"users" form:"users" example:"[550e8400-e29b-41d4-a716-446655440000,550e8400-e29b-41d4-a716-446655440000]" format:"[]uuid"`
}

type DeleteTask struct {
	Id string `form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}

type TasksByProject struct {
	Id string `form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}
