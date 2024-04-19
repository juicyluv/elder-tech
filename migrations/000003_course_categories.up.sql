create table if not exists course_categories
(
    id   smallserial primary key,
    name text not null unique
);

INSERT INTO course_categories(name)
VALUES ('Базовый'),
       ('Языки'),
       ('Программирование'),
       ('Психология'),
       ('Общее'),
       ('Вязание'),
       ('Шахматы'),
       ('Готовка'),
       ('Стройка и ремонт'),
       ('Компьютерная грамотность');
