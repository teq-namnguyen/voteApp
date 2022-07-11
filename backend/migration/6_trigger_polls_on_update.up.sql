create trigger update_users_on_update
    before update
    on
        polls
    for each row
execute procedure update_updated_on_user_task();
