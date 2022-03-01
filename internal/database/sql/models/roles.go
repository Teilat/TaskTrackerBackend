package models

import "github.com/google/uuid"

//go:generate reform

// Roles represents a row in Roles table.
//reform:Roles
type Roles struct {
	Id       uuid.UUID `reform:"Id,pk"` // FIXME unhandled database type "uniqueidentifier"
	RoleName string    `reform:"RoleName"`
}
