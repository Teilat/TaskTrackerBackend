package models

//go:generate reform

// Tags represents a row in Tags table.
//reform:Tags
type Tags struct {
	Id    int32  `reform:"Id,pk"`
	Name  string `reform:"Name"`
	Color string `reform:"Color"`
}
