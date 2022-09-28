CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  email VARCHAR(60) UNIQUE NOT NULL,
  password VARCHAR(50) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp
);

CREATE INDEX idx_email ON users(email)