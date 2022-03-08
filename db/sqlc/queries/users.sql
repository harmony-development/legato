-- name: CreateUser :one
INSERT INTO users (id, email, username, password_hash) VALUES (generate_id(), $1, $2, $3) RETURNING id;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, password_hash FROM users WHERE email = $1 LIMIT 1;