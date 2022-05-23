package models

import "github.com/google/uuid"

//go:generate reform

// ProjectAndUsers represents a row in ProjectAndUsers table.
//reform:ProjectAndUsers
type ProjectAndUsers struct {
	Id        uuid.UUID `reform:"Id,pk"`
	UserId    uuid.UUID `reform:"UserId"`
	ProjectId uuid.UUID `reform:"ProjectId"`
}
