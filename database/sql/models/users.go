package models

import (
	"github.com/google/uuid"
)

//go:generate reform

// Users represents a row in Users table.
//reform:Users
type Users struct {
	Id       uuid.UUID `reform:"Id,pk"`
	Name     string    `reform:"Name"`
	Surname  string    `reform:"Surname"`
	Nickname string    `reform:"Nickname"`
	RoleId   uuid.UUID `reform:"RoleId"`
	Password string    `reform:"Password"`
}
