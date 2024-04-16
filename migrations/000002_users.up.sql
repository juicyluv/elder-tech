create table if not exists users
(
    id           bigserial primary key,
    name         text        not null,
    phone        text        not null,
    password_enc text        not null,
    created_at   timestamptz not null,
    surname      text,
    patronymic   text,
    age          int2,
    gender       int2,
    email        text,
    image_id     int8 references images (id),
    last_online  timestamptz,
    deleted_at   timestamptz
);

INSERT INTO users (name, phone, password_enc, created_at, surname, patronymic, age, gender, email, image_id, last_online, deleted_at)
VALUES
    ('John', '123456789', 'pass123', NOW(), 'Doe', 'Michael', 30, 1, 'john@example.com', 1, NOW(), NULL),
    ('Alice', '987654321', 'pass123', NOW(), 'Smith', 'Anne', 25, 0, 'alice@example.com', 2, NOW(), NULL),
    ('Bob', '555555555', 'pass123', NOW(), 'Johnson', 'Robert', 40, 1, 'bob@example.com', 3, NOW(), NULL);
