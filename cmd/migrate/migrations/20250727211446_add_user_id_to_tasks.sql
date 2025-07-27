-- +goose Up
-- +goose StatementBegin
-- Delete existing tasks since they don't have users
DELETE FROM tasks;

-- Now add the column
ALTER TABLE tasks ADD COLUMN user_id TEXT NOT NULL;

-- Create index
CREATE INDEX idx_tasks_user_id ON tasks(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_tasks_user_id;
ALTER TABLE tasks DROP COLUMN user_id;
-- +goose StatementEnd