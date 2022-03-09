package models

import ()

//go:generate reform

// Tasks represents a row in Tasks table.
//reform:Tasks
type Tasks struct {
	Id              []byte  `reform:"Id,pk"`     // FIXME unhandled database type "uniqueidentifier"
	ProjectId       []byte  `reform:"ProjectId"` // FIXME unhandled database type "uniqueidentifier"
	TaskTitle       string  `reform:"TaskTitle"`
	TaskDescription *string `reform:"TaskDescription"`
}
