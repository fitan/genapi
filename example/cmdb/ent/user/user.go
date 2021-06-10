// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeRoleBindings holds the string denoting the role_bindings edge name in mutations.
	EdgeRoleBindings = "role_bindings"
	// EdgeAlert holds the string denoting the alert edge name in mutations.
	EdgeAlert = "alert"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RoleBindingsTable is the table the holds the role_bindings relation/edge.
	RoleBindingsTable = "role_bindings"
	// RoleBindingsInverseTable is the table name for the RoleBinding entity.
	// It exists in this package in order to avoid circular dependency with the "rolebinding" package.
	RoleBindingsInverseTable = "role_bindings"
	// RoleBindingsColumn is the table column denoting the role_bindings relation/edge.
	RoleBindingsColumn = "user_role_bindings"
	// AlertTable is the table the holds the alert relation/edge.
	AlertTable = "users"
	// AlertInverseTable is the table name for the Alert entity.
	// It exists in this package in order to avoid circular dependency with the "alert" package.
	AlertInverseTable = "alerts"
	// AlertColumn is the table column denoting the alert relation/edge.
	AlertColumn = "user_alert"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldPassword,
	FieldEmail,
	FieldPhone,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_alert",
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
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
	RoleTourist Role = "tourist"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUser, RoleAdmin, RoleTourist:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}
