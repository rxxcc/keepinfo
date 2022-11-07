ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT "sessions_pkey";

ALTER TABLE IF EXISTS "contacts" DROP CONSTRAINT "contacts_pkey";

DROP TABLE IF EXISTS "users";