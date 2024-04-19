create table if not exists course_progresses
(
    course_id int4 not null references courses (id)  ON DELETE CASCADE,
    user_id   int8 not null references users (id)  ON DELETE CASCADE,
    lesson_id int8 not null references course_block_lessons (id)  ON DELETE CASCADE
);
