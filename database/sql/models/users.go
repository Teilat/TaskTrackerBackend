package models

//go:generate reform

// Users represents a row in Users table.
//reform:Users
type Users struct {
	Id       int32  `reform:"Id,pk"`
	Name     string `reform:"Name"`
	Surname  string `reform:"Surname"`
	Nickname string `reform:"Nickname"`
	RoleId   int32  `reform:"RoleId"`
	Password string `reform:"Password"`
}
