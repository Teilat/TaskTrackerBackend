package models

import "github.com/google/uuid"

//go:generate reform

// Tasks represents a row in Tasks table.
//reform:Tasks
type Tasks struct {
	Id              uuid.UUID `reform:"Id,pk"`     // FIXME unhandled database type "uniqueidentifier"
	ProjectId       uuid.UUID `reform:"ProjectId"` // FIXME unhandled database type "uniqueidentifier"
	TaskTitle       string    `reform:"TaskTitle"`
	TaskDescription *string   `reform:"TaskDescription"`
}
