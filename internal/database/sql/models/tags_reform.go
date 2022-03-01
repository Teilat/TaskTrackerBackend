// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"github.com/google/uuid"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type tagsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *tagsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("Tags").
func (v *tagsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *tagsTableType) Columns() []string {
	return []string{
		"Id",
		"TagName",
		"TagColor",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *tagsTableType) NewStruct() reform.Struct {
	return new(Tags)
}

// NewRecord makes a new record for that table.
func (v *tagsTableType) NewRecord() reform.Record {
	return new(Tags)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *tagsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// TagsTable represents Tags view or table in SQL database.
var TagsTable = &tagsTableType{
	s: parse.StructInfo{
		Type:    "Tags",
		SQLName: "Tags",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "uuid.UUID", Column: "Id"},
			{Name: "TagName", Type: "string", Column: "TagName"},
			{Name: "TagColor", Type: "string", Column: "TagColor"},
		},
		PKFieldIndex: 0,
	},
	z: new(Tags).Values(),
}

// String returns a string representation of this struct or record.
func (s Tags) String() string {
	res := make([]string, 3)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "TagName: " + reform.Inspect(s.TagName, true)
	res[2] = "TagColor: " + reform.Inspect(s.TagColor, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Tags) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.TagName,
		s.TagColor,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Tags) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.TagName,
		&s.TagColor,
	}
}

// View returns View object for that struct.
func (s *Tags) View() reform.View {
	return TagsTable
}

// Table returns Table object for that record.
func (s *Tags) Table() reform.Table {
	return TagsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Tags) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Tags) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Tags) HasPK() bool {
	return s.Id != TagsTable.z[TagsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *Tags) SetPK(pk interface{}) {
	reform.SetPK(s, pk.(uuid.UUID))
}

// check interfaces
var (
	_ reform.View   = TagsTable
	_ reform.Struct = (*Tags)(nil)
	_ reform.Table  = TagsTable
	_ reform.Record = (*Tags)(nil)
	_ fmt.Stringer  = (*Tags)(nil)
)

func init() {
	parse.AssertUpToDate(&TagsTable.s, new(Tags))
}
