// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type rolesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *rolesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("Roles").
func (v *rolesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *rolesTableType) Columns() []string {
	return []string{
		"Id",
		"Name",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *rolesTableType) NewStruct() reform.Struct {
	return new(Roles)
}

// NewRecord makes a new record for that table.
func (v *rolesTableType) NewRecord() reform.Record {
	return new(Roles)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *rolesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// RolesTable represents Roles view or table in SQL database.
var RolesTable = &rolesTableType{
	s: parse.StructInfo{
		Type:    "Roles",
		SQLName: "Roles",
		Fields: []parse.FieldInfo{
			{Name: "Id", Type: "int32", Column: "Id"},
			{Name: "Name", Type: "string", Column: "Name"},
		},
		PKFieldIndex: 0,
	},
	z: new(Roles).Values(),
}

// String returns a string representation of this struct or record.
func (s Roles) String() string {
	res := make([]string, 2)
	res[0] = "Id: " + reform.Inspect(s.Id, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Roles) Values() []interface{} {
	return []interface{}{
		s.Id,
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Roles) Pointers() []interface{} {
	return []interface{}{
		&s.Id,
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *Roles) View() reform.View {
	return RolesTable
}

// Table returns Table object for that record.
func (s *Roles) Table() reform.Table {
	return RolesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Roles) PKValue() interface{} {
	return s.Id
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Roles) PKPointer() interface{} {
	return &s.Id
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Roles) HasPK() bool {
	return s.Id != RolesTable.z[RolesTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.Id = pk.
func (s *Roles) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = RolesTable
	_ reform.Struct = (*Roles)(nil)
	_ reform.Table  = RolesTable
	_ reform.Record = (*Roles)(nil)
	_ fmt.Stringer  = (*Roles)(nil)
)

func init() {
	parse.AssertUpToDate(&RolesTable.s, new(Roles))
}
