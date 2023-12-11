-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  chat_id INTEGER,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
