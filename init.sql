-- Create Database
-- CREATE DATABASE booksdb ENCODING = 'UTF8';

-- -- Create User
-- CREATE USER admin WITH ENCRYPTED PASSWORD 'admin';

GRANT CONNECT ON DATABASE booksdb TO admin;
GRANT USAGE ON SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;

\connect booksdb

--Creation
CREATE SCHEMA IF NOT EXISTS "booksdb";

GRANT ALL PRIVILEGES ON SCHEMA "booksdb" TO admin;
GRANT USAGE ON SCHEMA "booksdb" TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA "booksdb" TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA "booksdb" TO admin;

--set default schema
ALTER USER admin SET search_path = "booksdb";

SET search_path TO "booksdb";

CREATE OR REPLACE LANGUAGE plpgsql;