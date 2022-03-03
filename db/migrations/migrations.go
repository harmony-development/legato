// Package migrations contains the migration scripts for legato.
package migrations

import "github.com/harmony-development/legato/db"

type Migration func(db *db.DB) error

var Migrations = []Migration{
	Migration0,
}
