CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE contacts (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id uuid,
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

ALTER TABLE contacts  
    ADD CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
                ON DELETE CASCADE