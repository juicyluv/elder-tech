create table if not exists course_block_lesson_content
(
    course_block_lesson_id int8 not null references course_block_lessons (id)  ON DELETE CASCADE,
    type                   int2 not null,
    value                  text not null
);
