-- +goose Up
CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    name TEXT,
    race TEXT,
    class TEXT,
    level INT DEFAULT 1,
    str INT DEFAULT 10,
    dex INT DEFAULT 10,
    con INT DEFAULT 10,
    int INT DEFAULT 10,
    wis INT DEFAULT 10,
    cha INT DEFAULT 10,
    hp INT DEFAULT 10,
    ac INT DEFAULT 10,
    notes JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO characters (id, name, notes) VALUES (1, 'Zorya Yaroslava', '{}'::jsonb);


-- +goose Down
DROP TABLE characters;