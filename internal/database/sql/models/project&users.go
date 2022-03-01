package models

import "github.com/google/uuid"

//go:generate reform

// Project&Users represents a row in Project&Users table.
//reform:Project&Users
type ProjectAndUsers struct {
	Id        uuid.UUID `reform:"Id,pk"`     // FIXME unhandled database type "uniqueidentifier"
	UserId    uuid.UUID `reform:"UserId"`    // FIXME unhandled database type "uniqueidentifier"
	ProjectId uuid.UUID `reform:"ProjectId"` // FIXME unhandled database type "uniqueidentifier"
}
