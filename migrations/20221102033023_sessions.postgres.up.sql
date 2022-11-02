CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE sessions (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid,
    email VARCHAR UNIQUE NOT NULL,
    refresh_token VARCHAR NOT NULL,
    user_agent VARCHAR NOT NULL,
    client_ip VARCHAR NOT NULL,
    is_blocked BOOLEAN NOT NULL DEFAULT false,
    expired_at timestamp NOT NULL,
    created_at timestamp NOT NULL
);

ALTER TABLE sessions  
    ADD CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
                ON DELETE CASCADE