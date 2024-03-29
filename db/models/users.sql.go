// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package models

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, email, username, password_hash) VALUES (generate_id(), $1, $2, $3) RETURNING id
`

type CreateUserParams struct {
	Email        *string
	Username     string
	PasswordHash []byte
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uint64, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Username, arg.PasswordHash)
	var id uint64
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uint64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash FROM users WHERE email = $1 LIMIT 1
`

type GetUserByEmailRow struct {
	ID           uint64
	Email        *string
	PasswordHash []byte
}

func (q *Queries) GetUserByEmail(ctx context.Context, email *string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(&i.ID, &i.Email, &i.PasswordHash)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, avatar, user_status, host, remote_id, email, password_hash, created FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id uint64) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Avatar,
		&i.UserStatus,
		&i.Host,
		&i.RemoteID,
		&i.Email,
		&i.PasswordHash,
		&i.Created,
	)
	return i, err
}
