package models

//go:generate reform

// TaskAndUsers represents a row in TaskAndUsers table.
//reform:TaskAndUsers
type TaskAndUsers struct {
	Id     int32 `reform:"Id,pk"`
	TaskId int32 `reform:"TaskId"`
	UserId int32 `reform:"UserId"`
}
