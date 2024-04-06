create table if not exists course_categories
(
    id   smallserial primary key,
    name text not null unique
);