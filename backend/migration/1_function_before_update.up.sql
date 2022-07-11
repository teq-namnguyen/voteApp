CREATE FUNCTION update_updated_on_user_task()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language plpgsql;
