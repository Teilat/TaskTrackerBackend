package models

import (
	"time"
)

//go:generate reform

// Projects represents a row in Projects table.
//reform:Projects
type Projects struct {
	Id                 []byte    `reform:"Id,pk"`    // FIXME unhandled database type "uniqueidentifier"
	ParentId           []byte    `reform:"ParentId"` // FIXME unhandled database type "uniqueidentifier"
	ProjectName        string    `reform:"ProjectName"`
	ProjectDescription *string   `reform:"ProjectDescription"`
	CreationDate       time.Time `reform:"CreationDate"`
	OwnerId            []byte    `reform:"OwnerId"` // FIXME unhandled database type "uniqueidentifier"
}
