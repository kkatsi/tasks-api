-- +goose Up
-- Add unique constraints to existing users table
-- +goose StatementBegin
CREATE UNIQUE INDEX idx_users_email_unique ON users(email);
CREATE UNIQUE INDEX idx_users_username_unique ON users(username);
-- +goose StatementEnd

-- +goose Down
-- Remove unique constraints
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_email_unique;
DROP INDEX IF EXISTS idx_users_username_unique;
-- +goose StatementEnd
