CREATE TABLE users (
    id      serial PRIMARY KEY,
    name    text NOT NULL,
    email   text NOT NULL UNIQUE,
    picture text
);
