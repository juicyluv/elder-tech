create table if not exists course_blocks
(
    id          bigserial primary key,
    course_id   int4 not null references courses (id),
    number      int2 not null,
    title       text not null,
    description text not null
);