package models

import "github.com/google/uuid"

type Task struct {
	Id              uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ProjectId       uuid.UUID `json:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	TaskTitle       string    `json:"taskTitle" example:"task name"`
	TaskDescription string    `json:"taskDescription" example:"example task"`
}

type AddTask struct {
	ProjectId       uuid.UUID `json:"projectId" form:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	TaskTitle       string    `json:"taskTitle" form:"taskTitle" example:"task name"`
	TaskDescription string    `json:"taskDescription" form:"taskDescription" example:"example task"`
}

type UpdateTask struct {
	Id              uuid.UUID `json:"id" form:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	ProjectId       uuid.UUID `json:"projectId" form:"projectId" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	TaskTitle       string    `json:"taskTitle" form:"taskTitle" example:"task name"`
	TaskDescription string    `json:"taskDescription" form:"taskDescription" example:"example task"`
}

type DeleteTask struct {
	Id uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
}
