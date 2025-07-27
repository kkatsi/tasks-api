-- name: GetTask :one
SELECT * FROM tasks WHERE id = ? AND user_id = ?;

-- name: GetTasks :many
SELECT * FROM tasks WHERE user_id = ? ORDER BY tasks.created_at DESC LIMIT ? OFFSET ?;

-- name: CreateTask :one
INSERT INTO tasks (id, title, description, completed, created_at, updated_at, user_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteTask :one
DELETE FROM tasks WHERE id = ? AND user_id = ?
RETURNING id;

-- name: UpdateTask :one
UPDATE tasks
SET title = ?, description = ?, completed = ?, updated_at = ?
WHERE id = ? AND user_id = ?
RETURNING *;