package models

//go:generate reform

// Roles represents a row in Roles table.
//reform:Roles
type Roles struct {
	Id   int32  `reform:"Id,pk"`
	Name string `reform:"Name"`
}
