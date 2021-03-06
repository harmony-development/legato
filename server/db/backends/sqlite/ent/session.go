// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/harmony-development/legato/server/db/backends/sqlite/ent/localuser"
	"github.com/harmony-development/legato/server/db/backends/sqlite/ent/session"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SessionID holds the value of the "sessionID" field.
	SessionID string `json:"sessionID,omitempty"`
	// Expires holds the value of the "expires" field.
	Expires time.Time `json:"expires,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges               SessionEdges `json:"edges"`
	local_user_sessions *int
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// User holds the value of the user edge.
	User *LocalUser `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*LocalUser, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: localuser.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			values[i] = &sql.NullInt64{}
		case session.FieldSessionID:
			values[i] = &sql.NullString{}
		case session.FieldExpires:
			values[i] = &sql.NullTime{}
		case session.ForeignKeys[0]: // local_user_sessions
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldSessionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sessionID", values[i])
			} else if value.Valid {
				s.SessionID = value.String
			}
		case session.FieldExpires:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires", values[i])
			} else if value.Valid {
				s.Expires = value.Time
			}
		case session.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field local_user_sessions", value)
			} else if value.Valid {
				s.local_user_sessions = new(int)
				*s.local_user_sessions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Session entity.
func (s *Session) QueryUser() *LocalUserQuery {
	return (&SessionClient{config: s.config}).QueryUser(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", sessionID=")
	builder.WriteString(s.SessionID)
	builder.WriteString(", expires=")
	builder.WriteString(s.Expires.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
