CREATE TABLE IF NOT EXISTS short_url(
    id INTEGER PRIMARY KEY,
    src VARCHAR(255) UNIQUE NOT NULL DEFAULT ""
);
