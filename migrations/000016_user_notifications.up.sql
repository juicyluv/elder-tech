create table if not exists notifications
(
    id         bigserial primary key,
    user_id    int8        not null references users (id)  ON DELETE CASCADE,
    title      text        not null,
    content    text        not null,
    created_at timestamptz not null,
    seen       bool        not null
);
