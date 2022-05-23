package models

import (
)

//go:generate reform

// Users represents a row in Users table.
//reform:Users
type Users struct {
    Id []byte `reform:"Id,pk"` // FIXME unhandled database type "uniqueidentifier"
    Name string `reform:"Name"` 
    Surname string `reform:"Surname"` 
    Nickname string `reform:"Nickname"` 
    RoleId []byte `reform:"RoleId"` // FIXME unhandled database type "uniqueidentifier"
    Password string `reform:"Password"` 
}
