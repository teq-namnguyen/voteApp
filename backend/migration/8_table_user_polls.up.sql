create table if not exists user_polls
(
    id      serial unique not null,
    user_id int           not null,
    poll_id int           not null,
    unique (user_id, poll_id),
    constraint fk_users
        foreign key (user_id) references users (id),
    constraint fk_polls
        foreign key (poll_id) references polls (id)
)