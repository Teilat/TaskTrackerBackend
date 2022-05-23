package models

import (
)

//go:generate reform

// Columns represents a row in Columns table.
//reform:Columns
type Columns struct {
    ID []byte `reform:"id,pk"` // FIXME unhandled database type "uniqueidentifier"
    Name string `reform:"Name"` 
}
