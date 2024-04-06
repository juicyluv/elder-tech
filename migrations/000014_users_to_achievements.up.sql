create table if not exists users_to_achievements
(
    user_id        int8        not null references users (id),
    achievement_id int8        not null references achievements (id),
    created_at     timestamptz not null
);