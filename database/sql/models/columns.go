package models

import (
	"github.com/google/uuid"
)

//go:generate reform

// Columns represents a row in Columns table.
//reform:Columns
type Columns struct {
	ID   uuid.UUID `reform:"id,pk"`
	Name string    `reform:"Name"`
}
