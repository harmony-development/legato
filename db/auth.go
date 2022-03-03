package db

import (
	"context"
)

// SaveUser saves a user to the database.
func (db *DB) SaveUser(ctx context.Context, email string, username string, passwordHash []byte) error {
	_, err := db.conn.Exec(
		ctx,
		"INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3)",
		email, username, passwordHash,
	)

	return TryWrap(err, "failed to save user")
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*UserAccount, error) {
	account := UserAccount{}

	err := db.conn.
		QueryRow(ctx, "SELECT id, email, username, password_hash, created FROM users WHERE email = $1 LIMIT 1", email).
		Scan(&account.Id, &account.Email, &account.Username, &account.Password_Hash, &account.Created)

	return &account, TryWrap(err, "failed to get user by email")
}

func (db *DB) GetUserByID(ctx context.Context, id int64) (*User, error) {
	user := User{}

	err := db.conn.
		QueryRow(ctx, "SELECT id, username, created FROM users WHERE id = $1 LIMIT 1", id).
		Scan(&user.Id, &user.Username, &user.Created)

	return &user, TryWrap(err, "failed to get user by id")
}
