// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type sysdiagramsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *sysdiagramsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("sysdiagrams").
func (v *sysdiagramsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *sysdiagramsTableType) Columns() []string {
	return []string{
		"name",
		"principal_id",
		"diagram_id",
		"version",
		"definition",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *sysdiagramsTableType) NewStruct() reform.Struct {
	return new(Sysdiagrams)
}

// NewRecord makes a new record for that table.
func (v *sysdiagramsTableType) NewRecord() reform.Record {
	return new(Sysdiagrams)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *sysdiagramsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// SysdiagramsTable represents sysdiagrams view or table in SQL database.
var SysdiagramsTable = &sysdiagramsTableType{
	s: parse.StructInfo{
		Type:    "Sysdiagrams",
		SQLName: "sysdiagrams",
		Fields: []parse.FieldInfo{
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "PrincipalID", Type: "int32", Column: "principal_id"},
			{Name: "DiagramID", Type: "int32", Column: "diagram_id"},
			{Name: "Version", Type: "*int32", Column: "version"},
			{Name: "Definition", Type: "[]uint8", Column: "definition"},
		},
		PKFieldIndex: 2,
	},
	z: new(Sysdiagrams).Values(),
}

// String returns a string representation of this struct or record.
func (s Sysdiagrams) String() string {
	res := make([]string, 5)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "PrincipalID: " + reform.Inspect(s.PrincipalID, true)
	res[2] = "DiagramID: " + reform.Inspect(s.DiagramID, true)
	res[3] = "Version: " + reform.Inspect(s.Version, true)
	res[4] = "Definition: " + reform.Inspect(s.Definition, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Sysdiagrams) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.PrincipalID,
		s.DiagramID,
		s.Version,
		s.Definition,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Sysdiagrams) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.PrincipalID,
		&s.DiagramID,
		&s.Version,
		&s.Definition,
	}
}

// View returns View object for that struct.
func (s *Sysdiagrams) View() reform.View {
	return SysdiagramsTable
}

// Table returns Table object for that record.
func (s *Sysdiagrams) Table() reform.Table {
	return SysdiagramsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Sysdiagrams) PKValue() interface{} {
	return s.DiagramID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Sysdiagrams) PKPointer() interface{} {
	return &s.DiagramID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Sysdiagrams) HasPK() bool {
	return s.DiagramID != SysdiagramsTable.z[SysdiagramsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.DiagramID = pk.
func (s *Sysdiagrams) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = SysdiagramsTable
	_ reform.Struct = (*Sysdiagrams)(nil)
	_ reform.Table  = SysdiagramsTable
	_ reform.Record = (*Sysdiagrams)(nil)
	_ fmt.Stringer  = (*Sysdiagrams)(nil)
)

func init() {
	parse.AssertUpToDate(&SysdiagramsTable.s, new(Sysdiagrams))
}
