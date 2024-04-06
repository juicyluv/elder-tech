create table if not exists course_ratings
(
    course_id int4 not null references courses (id),
    user_id   int8 not null references users (id),
    rating    int2 not null
);