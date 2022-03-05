package middleware

import "github.com/harmony-development/legato/db"

type MiddlewareHandler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *MiddlewareHandler {
	return &MiddlewareHandler{
		db: db,
	}
}