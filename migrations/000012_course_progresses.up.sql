create table if not exists course_progresses
(
    course_id int4 not null references courses (id),
    user_id   int8 not null references users (id),
    lesson_id int8 not null references course_block_lessons (id)
);