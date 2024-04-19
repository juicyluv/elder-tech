create table if not exists favourite_courses
(
    course_id int4 not null references courses (id)  ON DELETE CASCADE,
    user_id   int8 not null references users (id)  ON DELETE CASCADE
);
