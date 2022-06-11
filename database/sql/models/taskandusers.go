package models

import "github.com/google/uuid"

//go:generate reform

// TaskAndUsers represents a row in TaskAndUsers table.
//reform:TaskAndUsers
type TaskAndUsers struct {
	Id     uuid.UUID `reform:"Id,pk"`
	TaskId uuid.UUID `reform:"TaskId"`
	UserId uuid.UUID `reform:"UserId"`
}
