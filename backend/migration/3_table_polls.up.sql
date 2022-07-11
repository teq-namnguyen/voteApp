create table if not exists polls
(
    id         serial unique not null,
    user_id    int           not null,
    question   varchar(255)  not null,
    created_at timestamp     not null default current_timestamp,
    updated_at timestamp     not null default current_timestamp
);
