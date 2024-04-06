create table if not exists course_block_lessons
(
    id              bigserial primary key,
    course_block_id int8 not null references course_blocks (id),
    number          int2 not null,
    title           text not null,
    description     text not null
);