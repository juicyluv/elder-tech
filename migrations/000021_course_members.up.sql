create table if not exists course_members (
    course_id int4 not null references courses(id)  ON DELETE CASCADE,
    user_id int8 not null references users(id)
);
