// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type columnsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *columnsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("Columns").
func (v *columnsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *columnsTableType) Columns() []string {
	return []string{
		"Id",
		"Name",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *columnsTableType) NewStruct() reform.Struct {
	return new(Columns)
}

// NewRecord makes a new record for that table.
func (v *columnsTableType) NewRecord() reform.Record {
	return new(Columns)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *columnsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ColumnsTable represents Columns view or table in SQL database.
var ColumnsTable = &columnsTableType{
	s: parse.StructInfo{
		Type:    "Columns",
		SQLName: "Columns",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "int32", Column: "Id"},
			{Name: "Name", Type: "string", Column: "Name"},
		},
		PKFieldIndex: 0,
	},
	z: new(Columns).Values(),
}

// String returns a string representation of this struct or record.
func (s Columns) String() string {
	res := make([]string, 2)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Columns) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Columns) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *Columns) View() reform.View {
	return ColumnsTable
}

// Table returns Table object for that record.
func (s *Columns) Table() reform.Table {
	return ColumnsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Columns) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Columns) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Columns) HasPK() bool {
	return s.Id != ColumnsTable.z[ColumnsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *Columns) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = ColumnsTable
	_ reform.Struct = (*Columns)(nil)
	_ reform.Table  = ColumnsTable
	_ reform.Record = (*Columns)(nil)
	_ fmt.Stringer  = (*Columns)(nil)
)

func init() {
	parse.AssertUpToDate(&ColumnsTable.s, new(Columns))
}
