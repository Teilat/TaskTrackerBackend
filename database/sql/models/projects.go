package models

import (
	"github.com/google/uuid"
	"time"
)

//go:generate reform

// Projects represents a row in Projects table.
//reform:Projects
type Projects struct {
	Id           uuid.UUID  `reform:"Id,pk"`
	ParentId     *uuid.UUID `reform:"ParentId"`
	Name         string     `reform:"Name"`
	Description  *string    `reform:"Description"`
	CreationDate time.Time  `reform:"CreationDate"`
	OwnerId      uuid.UUID  `reform:"OwnerId"`
}
