// Code generated by entc, DO NOT EDIT.

package rolebinding

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the rolebinding type in the database.
	Label = "role_binding"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"

	// Table holds the table name of the rolebinding in the database.
	Table = "role_bindings"
)

// Columns holds all SQL columns for rolebinding fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the RoleBinding type.
var ForeignKeys = []string{
	"project_role_bindings",
	"service_role_bindings",
	"user_role_bindings",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleUser:
		return nil
	default:
		return fmt.Errorf("rolebinding: invalid enum value for role field: %q", r)
	}
}
