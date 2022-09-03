-- +goose Up
-- +goose StatementBegin
CREATE TABLE contacts (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  phone VARCHAR(50) NOT NULL,
  label TEXT [] NOT NULL,
  address VARCHAR(500) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE contacts;
-- +goose StatementEnd
