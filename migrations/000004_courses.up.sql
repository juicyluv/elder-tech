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