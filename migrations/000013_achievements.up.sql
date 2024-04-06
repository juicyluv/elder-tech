create table if not exists achievements
(
    id          smallserial primary key,
    title       text not null,
    description text not null,
    image_id    int8 references images (id)
);