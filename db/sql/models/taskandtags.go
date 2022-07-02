package models

//go:generate reform

// TaskAndTags represents a row in TaskAndTags table.
//reform:TaskAndTags
type TaskAndTags struct {
	Id     int32 `reform:"Id,pk"`
	TaskId int32 `reform:"TaskId"`
	TagId  int32 `reform:"TagId"`
}
