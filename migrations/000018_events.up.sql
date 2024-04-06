create table if not exists events
(
    id          bigserial primary key,
    entity_id   int8        not null,
    entity_type int2        not null,
    time        timestamptz not null,
    details     text        not null
);