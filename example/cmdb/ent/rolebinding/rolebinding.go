// Code generated by entc, DO NOT EDIT.

package rolebinding

import (
	"time"
)

const (
	// Label holds the string label denoting the rolebinding type in the database.
	Label = "role_binding"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleName holds the string denoting the role_name field in the database.
	FieldRoleName = "role_name"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// Table holds the table name of the rolebinding in the database.
	Table = "role_bindings"
)

// Columns holds all SQL columns for rolebinding fields.
var Columns = []string{
	FieldID,
	FieldRoleName,
	FieldRoleID,
	FieldStatus,
	FieldCreatedAt,
	FieldNote,
	FieldPermissions,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "role_bindings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_role_bind",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
