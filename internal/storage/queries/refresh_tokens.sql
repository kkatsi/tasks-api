-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE token_hash = ? AND expires_at > CURRENT_TIMESTAMP;

-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at)
VALUES (?, ?, ?, ?);

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens WHERE token_hash = ? AND user_id = ?;

-- name: DeleteExpiredTokens :exec
DELETE FROM refresh_tokens WHERE expires_at < CURRENT_TIMESTAMP;
