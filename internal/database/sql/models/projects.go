package models

import (
	"github.com/google/uuid"
	"time"
)

//go:generate reform

// Projects represents a row in Projects table.
//reform:Projects
type Projects struct {
	Id                 uuid.UUID `reform:"Id,pk"`    // FIXME unhandled database type "uniqueidentifier"
	ParentId           uuid.UUID `reform:"ParentId"` // FIXME unhandled database type "uniqueidentifier"
	ProjectName        string    `reform:"ProjectName"`
	ProjectDescription *string   `reform:"ProjectDescription"`
	CreationDate       time.Time `reform:"CreationDate"`
	OwnerId            uuid.UUID `reform:"OwnerId"` // FIXME unhandled database type "uniqueidentifier"
}
