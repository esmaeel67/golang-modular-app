-- +goose Up
-- Create the schema
CREATE SCHEMA IF NOT EXISTS mallbots;
SET SEARCH_PATH TO mallbots, public;


-- Create tables
-- CREATE TABLE stores_cache (
--     id         text        NOT NULL,
--     name       text        NOT NULL,
--     created_at timestamptz NOT NULL DEFAULT NOW(),
--     updated_at timestamptz NOT NULL DEFAULT NOW(),
--     PRIMARY KEY (id)
-- );


-- +goose Down
-- DROP TABLE stores_cache;