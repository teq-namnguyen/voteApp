create table if not exists options
(
    id         serial unique not null,
    poll_id    int           not null,
    name       varchar(128)  not null,
    vote       int           not null default 0,
    created_at timestamp     not null default current_timestamp,
    updated_at timestamp     not null default current_timestamp,
    constraint fk_polls
        foreign key (poll_id) references polls (id)
);
