package models

//go:generate reform

// Tasks represents a row in Tasks table.
//reform:Tasks
type Tasks struct {
	Id          int32   `reform:"Id,pk"`
	ProjectId   int32   `reform:"ProjectId"`
	Title       string  `reform:"Title"`
	Description *string `reform:"Description"`
	ColumnId    int32   `reform:"ColumnId"`
}
