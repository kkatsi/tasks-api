-- name: GetTask :one
SELECT * FROM tasks WHERE id = ?;

-- name: GetTasks :many
SELECT * FROM tasks ORDER BY tasks.created_at DESC;

-- name: CreateTask :one
INSERT INTO tasks (id, title, description, completed, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteTask :one
DELETE FROM tasks WHERE id = ?
RETURNING id;

-- name: UpdateTask :one
UPDATE tasks
SET title = ?, description = ?, completed = ?, updated_at = ?
WHERE id = ?
RETURNING *