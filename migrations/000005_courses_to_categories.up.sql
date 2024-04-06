create table if not exists courses_to_categories
(
    course_id          int4 not null references courses (id),
    course_category_id int2 not null references course_categories (id)
);