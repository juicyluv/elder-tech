alter table course_block_lesson_content add column number int2 not null default 1;
alter table course_block_lesson_content alter column number drop default;

alter table course_block_lesson_content add column id bigserial primary key;
