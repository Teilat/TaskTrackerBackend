package models

//go:generate reform

// Columns represents a row in Columns table.
//reform:Columns
type Columns struct {
	Id   int32  `reform:"Id,pk"`
	Name string `reform:"Name"`
}
