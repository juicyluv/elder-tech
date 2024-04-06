create table if not exists course_quizzes_questions
(
    quiz_id  int4 not null references course_quizzes (id),
    text     text not null,
    is_right bool not null
);