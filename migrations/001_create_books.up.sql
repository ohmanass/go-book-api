-- DB creation
CREATE TABLE books (
    id UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    title      TEXT NOT NULL,
    author     TEXT NOT NULL,
    year       INT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)