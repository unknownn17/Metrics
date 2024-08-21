CREATE TABLE IF NOT EXISTS books(
    id SERIAL PRIMARY KEY,
    title TEXT,
    author TEXT,
    published_year BIGINT
);