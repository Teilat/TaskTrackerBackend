package models

import "github.com/google/uuid"

//go:generate reform

// Tags represents a row in Tags table.
//reform:Tags
type Tags struct {
	Id       uuid.UUID `reform:"Id,pk"`
	TagName  string    `reform:"TagName"`
	TagColor string    `reform:"TagColor"`
}
