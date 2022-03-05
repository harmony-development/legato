package api

import "context"

type LegatoContext struct {
	context.Context
	UserID uint64
}
