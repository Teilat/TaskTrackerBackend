package models

//go:generate reform

// TaskAndTags represents a row in TaskAndTags table.
//reform:TaskAndTags
type TaskAndTags struct {
	Id     []byte `reform:"Id,pk"`  // FIXME unhandled database type "uniqueidentifier"
	TaskId []byte `reform:"TaskId"` // FIXME unhandled database type "uniqueidentifier"
	TagId  []byte `reform:"TagId"`  // FIXME unhandled database type "uniqueidentifier"
}
