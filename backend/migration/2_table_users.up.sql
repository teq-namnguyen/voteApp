CREATE TABLE IF NOT EXISTS users
(
    id         serial unique      not null,
    username   varchar(20) unique not null,
    password   varchar(128)       not null,
    created_at timestamp          not null default current_timestamp,
    updated_at timestamp          not null default current_timestamp
);
