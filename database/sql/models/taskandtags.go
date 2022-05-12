package models

import "github.com/google/uuid"

//go:generate reform

// TaskAndTags represents a row in TaskAndTags table.
//reform:TaskAndTags
type TaskAndTags struct {
	Id     uuid.UUID `reform:"Id,pk"`
	TaskId uuid.UUID `reform:"TaskId"`
	TagId  uuid.UUID `reform:"TagId"`
}
