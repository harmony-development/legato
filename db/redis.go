package db

import "strings"

const (
	authStepPrefix = "auth_state"
	sessionsPrefix = "sessions"
	chatPrefix     = "chat"
)

func subkey(key string, subkeys ...string) string {
	return key + ":" + strings.Join(subkeys, ":")
}
