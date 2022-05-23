package models

import (
)

//go:generate reform

// ProjectAndUsers represents a row in ProjectAndUsers table.
//reform:ProjectAndUsers
type ProjectAndUsers struct {
    Id []byte `reform:"Id,pk"` // FIXME unhandled database type "uniqueidentifier"
    UserId []byte `reform:"UserId"` // FIXME unhandled database type "uniqueidentifier"
    ProjectId []byte `reform:"ProjectId"` // FIXME unhandled database type "uniqueidentifier"
}
