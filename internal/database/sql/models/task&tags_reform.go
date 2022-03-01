// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type taskAndTagsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *taskAndTagsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("Task").
func (v *taskAndTagsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *taskAndTagsTableType) Columns() []string {
	return []string{
		"Id",
		"TaskId",
		"TagId",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *taskAndTagsTableType) NewStruct() reform.Struct {
	return new(TaskAndTags)
}

// NewRecord makes a new record for that table.
func (v *taskAndTagsTableType) NewRecord() reform.Record {
	return new(TaskAndTags)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *taskAndTagsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// TaskAndTagsTable represents Task view or table in SQL database.
var TaskAndTagsTable = &taskAndTagsTableType{
	s: parse.StructInfo{
		Type:    "TaskAndTags",
		SQLName: "Task",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "uuid.UUID", Column: "Id"},
			{Name: "TaskId", Type: "uuid.UUID", Column: "TaskId"},
			{Name: "TagId", Type: "uuid.UUID", Column: "TagId"},
		},
		PKFieldIndex: 0,
	},
	z: new(TaskAndTags).Values(),
}

// String returns a string representation of this struct or record.
func (s TaskAndTags) String() string {
	res := make([]string, 3)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "TaskId: " + reform.Inspect(s.TaskId, true)
	res[2] = "TagId: " + reform.Inspect(s.TagId, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *TaskAndTags) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.TaskId,
		s.TagId,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *TaskAndTags) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.TaskId,
		&s.TagId,
	}
}

// View returns View object for that struct.
func (s *TaskAndTags) View() reform.View {
	return TaskAndTagsTable
}

// Table returns Table object for that record.
func (s *TaskAndTags) Table() reform.Table {
	return TaskAndTagsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *TaskAndTags) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *TaskAndTags) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *TaskAndTags) HasPK() bool {
	return s.Id != TaskAndTagsTable.z[TaskAndTagsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *TaskAndTags) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = TaskAndTagsTable
	_ reform.Struct = (*TaskAndTags)(nil)
	_ reform.Table  = TaskAndTagsTable
	_ reform.Record = (*TaskAndTags)(nil)
	_ fmt.Stringer  = (*TaskAndTags)(nil)
)

func init() {
	parse.AssertUpToDate(&TaskAndTagsTable.s, new(TaskAndTags))
}
