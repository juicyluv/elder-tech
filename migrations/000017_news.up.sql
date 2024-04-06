create table if not exists news
(
    id         bigserial primary key,
    title      text        not null,
    content    text        not null,
    created_at timestamptz not null
);