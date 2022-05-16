package models

//go:generate reform

// Roles represents a row in Roles table.
//reform:Roles
type Roles struct {
	Id       []byte `reform:"Id,pk"` // FIXME unhandled database type "uniqueidentifier"
	RoleName string `reform:"RoleName"`
}
