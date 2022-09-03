-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  access_level INTEGER DEFAULT 1 NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
