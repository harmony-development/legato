// Code generated by entc, DO NOT EDIT.

package session

import (
	"time"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSessionID holds the string denoting the sessionid field in the database.
	FieldSessionID = "session_id"
	// FieldExpires holds the string denoting the expires field in the database.
	FieldExpires = "expires"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the session in the database.
	Table = "sessions"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "sessions"
	// UserInverseTable is the table name for the LocalUser entity.
	// It exists in this package in order to avoid circular dependency with the "localuser" package.
	UserInverseTable = "local_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "local_user_sessions"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldSessionID,
	FieldExpires,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Session type.
var ForeignKeys = []string{
	"local_user_sessions",
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
	// DefaultExpires holds the default value on creation for the "expires" field.
	DefaultExpires func() time.Time
)
