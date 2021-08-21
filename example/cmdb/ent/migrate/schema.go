// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AlertsColumns holds the columns for the "alerts" table.
	AlertsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// AlertsTable holds the schema information for the "alerts" table.
	AlertsTable = &schema.Table{
		Name:       "alerts",
		Columns:    AlertsColumns,
		PrimaryKey: []*schema.Column{AlertsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message", Type: field.TypeString},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
	}
	// RoleBindingsColumns holds the columns for the "role_bindings" table.
	RoleBindingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_name", Type: field.TypeString, Unique: true},
		{Name: "role_id", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "note", Type: field.TypeString, Size: 2147483647},
		{Name: "permissions", Type: field.TypeJSON},
		{Name: "user_role_bind", Type: field.TypeInt, Nullable: true},
	}
	// RoleBindingsTable holds the schema information for the "role_bindings" table.
	RoleBindingsTable = &schema.Table{
		Name:       "role_bindings",
		Columns:    RoleBindingsColumns,
		PrimaryKey: []*schema.Column{RoleBindingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_bindings_users_role_bind",
				Columns:    []*schema.Column{RoleBindingsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ServersColumns holds the columns for the "servers" table.
	ServersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ip", Type: field.TypeString, Unique: true},
		{Name: "machine_type", Type: field.TypeEnum, Enums: []string{"physical", "virtual"}},
		{Name: "platform_type", Type: field.TypeEnum, Enums: []string{"zstack", "k8s", "openstack"}},
		{Name: "system_type", Type: field.TypeEnum, Enums: []string{"linux", "windows"}},
		{Name: "service_tree_servers", Type: field.TypeInt, Nullable: true},
	}
	// ServersTable holds the schema information for the "servers" table.
	ServersTable = &schema.Table{
		Name:       "servers",
		Columns:    ServersColumns,
		PrimaryKey: []*schema.Column{ServersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "servers_service_trees_servers",
				Columns:    []*schema.Column{ServersColumns[5]},
				RefColumns: []*schema.Column{ServiceTreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ServiceTreesColumns holds the columns for the "service_trees" table.
	ServiceTreesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "note", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"project", "service"}},
		{Name: "service_tree_service", Type: field.TypeInt, Nullable: true},
	}
	// ServiceTreesTable holds the schema information for the "service_trees" table.
	ServiceTreesTable = &schema.Table{
		Name:       "service_trees",
		Columns:    ServiceTreesColumns,
		PrimaryKey: []*schema.Column{ServiceTreesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "service_trees_service_trees_service",
				Columns:    []*schema.Column{ServiceTreesColumns[4]},
				RefColumns: []*schema.Column{ServiceTreesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"user", "admin", "tourist"}},
		{Name: "user_msg", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_messages_msg",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserAlertColumns holds the columns for the "user_alert" table.
	UserAlertColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "alert_id", Type: field.TypeInt},
	}
	// UserAlertTable holds the schema information for the "user_alert" table.
	UserAlertTable = &schema.Table{
		Name:       "user_alert",
		Columns:    UserAlertColumns,
		PrimaryKey: []*schema.Column{UserAlertColumns[0], UserAlertColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_alert_user_id",
				Columns:    []*schema.Column{UserAlertColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_alert_alert_id",
				Columns:    []*schema.Column{UserAlertColumns[1]},
				RefColumns: []*schema.Column{AlertsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AlertsTable,
		MessagesTable,
		RoleBindingsTable,
		ServersTable,
		ServiceTreesTable,
		UsersTable,
		UserAlertTable,
	}
)

func init() {
	RoleBindingsTable.ForeignKeys[0].RefTable = UsersTable
	ServersTable.ForeignKeys[0].RefTable = ServiceTreesTable
	ServiceTreesTable.ForeignKeys[0].RefTable = ServiceTreesTable
	UsersTable.ForeignKeys[0].RefTable = MessagesTable
	UserAlertTable.ForeignKeys[0].RefTable = UsersTable
	UserAlertTable.ForeignKeys[1].RefTable = AlertsTable
}
