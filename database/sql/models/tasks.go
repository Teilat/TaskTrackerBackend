package models

import "github.com/google/uuid"

//go:generate reform

// Tasks represents a row in Tasks table.
//reform:Tasks
type Tasks struct {
	Id          uuid.UUID `reform:"Id,pk"`
	ProjectId   uuid.UUID `reform:"ProjectId"`
	Title       string    `reform:"Title"`
	Description *string   `reform:"Description"`
	ColumnId    uuid.UUID `reform:"ColumnId"`
}
