// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type projectAndUsersTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *projectAndUsersTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("ProjectAndUsers").
func (v *projectAndUsersTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *projectAndUsersTableType) Columns() []string {
	return []string{
		"Id",
		"ProjectId",
		"UserId",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *projectAndUsersTableType) NewStruct() reform.Struct {
	return new(ProjectAndUsers)
}

// NewRecord makes a new record for that table.
func (v *projectAndUsersTableType) NewRecord() reform.Record {
	return new(ProjectAndUsers)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *projectAndUsersTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ProjectAndUsersTable represents ProjectAndUsers view or table in SQL database.
var ProjectAndUsersTable = &projectAndUsersTableType{
	s: parse.StructInfo{
		Type:    "ProjectAndUsers",
		SQLName: "ProjectAndUsers",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "int32", Column: "Id"},
			{Name: "ProjectId", Type: "int32", Column: "ProjectId"},
			{Name: "UserId", Type: "int32", Column: "UserId"},
		},
		PKFieldIndex: 0,
	},
	z: new(ProjectAndUsers).Values(),
}

// String returns a string representation of this struct or record.
func (s ProjectAndUsers) String() string {
	res := make([]string, 3)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "ProjectId: " + reform.Inspect(s.ProjectId, true)
	res[2] = "UserId: " + reform.Inspect(s.UserId, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *ProjectAndUsers) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.ProjectId,
		s.UserId,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *ProjectAndUsers) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.ProjectId,
		&s.UserId,
	}
}

// View returns View object for that struct.
func (s *ProjectAndUsers) View() reform.View {
	return ProjectAndUsersTable
}

// Table returns Table object for that record.
func (s *ProjectAndUsers) Table() reform.Table {
	return ProjectAndUsersTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *ProjectAndUsers) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *ProjectAndUsers) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *ProjectAndUsers) HasPK() bool {
	return s.Id != ProjectAndUsersTable.z[ProjectAndUsersTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *ProjectAndUsers) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ProjectAndUsersTable
	_ reform.Struct = (*ProjectAndUsers)(nil)
	_ reform.Table  = ProjectAndUsersTable
	_ reform.Record = (*ProjectAndUsers)(nil)
	_ fmt.Stringer  = (*ProjectAndUsers)(nil)
)

func init() {
	parse.AssertUpToDate(&ProjectAndUsersTable.s, new(ProjectAndUsers))
}
