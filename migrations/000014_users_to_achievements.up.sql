create table if not exists users_to_achievements
(
    user_id        int8        not null references users (id)  ON DELETE CASCADE,
    achievement_id int8        not null references achievements (id)  ON DELETE CASCADE,
    created_at     timestamptz not null
);
