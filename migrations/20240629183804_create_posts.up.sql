CREATE TABLE IF NOT EXISTS Posts(
    id serial PRIMARY KEY,
    author varchar(64),
    title varchar(200),
    text varchar(3000),
    isCommented boolean
)