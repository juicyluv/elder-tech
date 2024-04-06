create table if not exists documents
(
    id       bigserial primary key,
    filename text not null,
    mime     int2 not null
);