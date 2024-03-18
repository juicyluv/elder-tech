create table if not exists user_images
(
    id       bigserial primary key,
    filename text not null
);

create table if not exists users
(
    id           bigserial primary key,
    type         int2        not null,
    name         text        not null,
    phone        text        not null,
    password_enc text        not null,
    created_at   timestamptz not null,
    surname      text,
    patronymic   text,
    age          int2,
    gender       int2,
    email        text,
    image_id     int8 references user_images (id),
    last_online  timestamptz,
    deleted_at   timestamptz
);

