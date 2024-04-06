create table if not exists course_quizzes
(
    id               serial primary key,
    course_lesson_id int4 not null references course_block_lessons (id),
    title            text not null
);