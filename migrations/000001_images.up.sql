create table if not exists images
(
    id       bigserial primary key,
    filename text not null
);