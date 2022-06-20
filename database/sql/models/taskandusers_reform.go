// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type taskAndUsersTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *taskAndUsersTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("TaskAndUsers").
func (v *taskAndUsersTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *taskAndUsersTableType) Columns() []string {
	return []string{
		"Id",
		"TaskId",
		"UserId",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *taskAndUsersTableType) NewStruct() reform.Struct {
	return new(TaskAndUsers)
}

// NewRecord makes a new record for that table.
func (v *taskAndUsersTableType) NewRecord() reform.Record {
	return new(TaskAndUsers)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *taskAndUsersTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// TaskAndUsersTable represents TaskAndUsers view or table in SQL database.
var TaskAndUsersTable = &taskAndUsersTableType{
	s: parse.StructInfo{
		Type:    "TaskAndUsers",
		SQLName: "TaskAndUsers",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "int32", Column: "Id"},
			{Name: "TaskId", Type: "int32", Column: "TaskId"},
			{Name: "UserId", Type: "int32", Column: "UserId"},
		},
		PKFieldIndex: 0,
	},
	z: new(TaskAndUsers).Values(),
}

// String returns a string representation of this struct or record.
func (s TaskAndUsers) String() string {
	res := make([]string, 3)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "TaskId: " + reform.Inspect(s.TaskId, true)
	res[2] = "UserId: " + reform.Inspect(s.UserId, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *TaskAndUsers) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.TaskId,
		s.UserId,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *TaskAndUsers) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.TaskId,
		&s.UserId,
	}
}

// View returns View object for that struct.
func (s *TaskAndUsers) View() reform.View {
	return TaskAndUsersTable
}

// Table returns Table object for that record.
func (s *TaskAndUsers) Table() reform.Table {
	return TaskAndUsersTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *TaskAndUsers) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *TaskAndUsers) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *TaskAndUsers) HasPK() bool {
	return s.Id != TaskAndUsersTable.z[TaskAndUsersTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *TaskAndUsers) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = TaskAndUsersTable
	_ reform.Struct = (*TaskAndUsers)(nil)
	_ reform.Table  = TaskAndUsersTable
	_ reform.Record = (*TaskAndUsers)(nil)
	_ fmt.Stringer  = (*TaskAndUsers)(nil)
)

func init() {
	parse.AssertUpToDate(&TaskAndUsersTable.s, new(TaskAndUsers))
}
