package models

import "github.com/google/uuid"

//go:generate reform

// Task&Tags represents a row in Task&Tags table.
//reform:Task&Tags
type TaskAndTags struct {
	Id     uuid.UUID `reform:"Id,pk"`  // FIXME unhandled database type "uniqueidentifier"
	TaskId uuid.UUID `reform:"TaskId"` // FIXME unhandled database type "uniqueidentifier"
	TagId  uuid.UUID `reform:"TagId"`  // FIXME unhandled database type "uniqueidentifier"
}
