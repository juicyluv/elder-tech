create table if not exists user_images
(
    id       uuid primary key,
    filename text not null
);

create table if not exists users
(
    id           uuid primary key,
    name         text        not null,
    surname      text        not null,
    patronymic   text        not null,
    age          int2        not null,
    gender       int2        not null,
    image_id     int2        not null,
    phone        text        not null,
    email        text        not null,
    password_enc text        not null,
    created_at   timestamptz not null,
    last_online  timestamp,
    deleted_at   timestamptz
);

