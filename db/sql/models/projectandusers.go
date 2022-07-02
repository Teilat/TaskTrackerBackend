package models

//go:generate reform

// ProjectAndUsers represents a row in ProjectAndUsers table.
//reform:ProjectAndUsers
type ProjectAndUsers struct {
	Id        int32 `reform:"Id,pk"`
	ProjectId int32 `reform:"ProjectId"`
	UserId    int32 `reform:"UserId"`
}
