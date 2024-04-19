create table if not exists courses
(
    id                       serial primary key,
    title                    text        not null,
    description              text        not null,
    author_id                int8        not null references users (id),
    difficulty               int2        not null,
    time_to_complete_minutes int4        not null,
    about                    text        not null,
    for_who                  text        not null,
    requirements             text        not null,
    created_at               timestamptz not null,
    updated_at               timestamptz,
    cover_image              int8 references images (id)
);

INSERT INTO courses (title, description, author_id, difficulty, time_to_complete_minutes, about, for_who, requirements, created_at, updated_at)
VALUES
    ('Introduction to Python Programming', 'Learn the basics of Python programming language.', 1, 1, 180, 'This course covers fundamental concepts such as variables, data types, loops, and functions.', 'Beginners interested in learning programming.', 'No prior programming experience required.', NOW(), NOW()),
    ('Intermediate JavaScript', 'Take your JavaScript skills to the next level.', 2, 2, 240, 'This course dives deeper into JavaScript with topics like object-oriented programming, asynchronous programming, and advanced DOM manipulation.', 'Those with basic knowledge of JavaScript.', 'Basic understanding of HTML, CSS, and JavaScript required.', NOW(), NOW()),
    ('Data Science Fundamentals', 'Explore the world of data science with Python.', 3, 3, 360, 'This course introduces key concepts in data science such as data manipulation, visualization, and machine learning algorithms.', 'Individuals interested in data analysis and machine learning.', 'Basic knowledge of Python recommended.', NOW(), NOW());
