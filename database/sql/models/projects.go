package models

import (
	"time"
)

//go:generate reform

// Projects represents a row in Projects table.
//reform:Projects
type Projects struct {
	Id           int32     `reform:"Id,pk"`
	ParentId     *int32    `reform:"ParentId"`
	Name         string    `reform:"Name"`
	Description  *string   `reform:"Description"`
	CreationDate time.Time `reform:"CreationDate"`
	OwnerId      int32     `reform:"OwnerId"`
}
