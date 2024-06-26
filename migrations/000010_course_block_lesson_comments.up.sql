create table if not exists course_block_lesson_comments
(
    id        bigserial primary key,
    author_id int8        not null references users (id)  ON DELETE CASCADE,
    time      timestamptz not null,
    comment   text
);
