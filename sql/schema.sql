create table if not exists images
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
    image_id     int8 references images (id),
    last_online  timestamptz,
    deleted_at   timestamptz
);

create table if not exists course_categories
(
    id   smallserial primary key,
    name text not null unique
);

create table if not exists courses
(
    id                       serial primary key,
    title                    text        not null,
    description              text        not null,
    author_id                int8        not null references users (id),
    difficulty               int2        not null,
    time_to_complete_minutes int4        not null,
    about                    text        not null,
    for_who                  text        not null,
    requirements             text        not null,
    created_at               timestamptz not null,
    updated_at               timestamptz,
    cover_image              int8 references images (id)
);

create table if not exists courses_to_categories
(
    course_id          int4 not null references courses (id),
    course_category_id int2 not null references course_categories (id)
);

create table if not exists course_ratings
(
    course_id int4 not null references courses (id),
    user_id   int8 not null references users (id),
    rating    int2 not null
);

create table if not exists course_blocks
(
    id          bigserial primary key,
    course_id   int4 not null references courses (id),
    number      int2 not null,
    title       text not null,
    description text not null
);

create table if not exists course_block_lessons
(
    id              bigserial primary key,
    course_block_id int8 not null references course_blocks (id),
    number          int2 not null,
    title           text not null,
    description     text not null
);

create table if not exists course_block_lesson_content
(
    course_block_lesson_id int8 not null references course_block_lessons (id),
    type                   int2 not null,
    value                  text not null
);

create table if not exists course_block_lesson_comments
(
    id        bigserial primary key,
    author_id int8        not null references users (id),
    time      timestamptz not null,
    comment   text
);

create table if not exists documents
(
    id       bigserial primary key,
    filename text not null,
    mime     int2 not null
);

create table if not exists course_progresses
(
    course_id int4 not null references courses (id),
    user_id   int8 not null references users (id),
    lesson_id int8 not null references course_block_lessons (id)
);