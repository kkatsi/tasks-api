-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :one
INSERT INTO users (id, username, email, password_hash, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = ?
RETURNING id;

-- name: UpdateUser :one
UPDATE users
SET email = ?, username = ?, updated_at = ?
WHERE id = ?
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET password_hash = ?, updated_at = ?
WHERE id = ?
RETURNING *;