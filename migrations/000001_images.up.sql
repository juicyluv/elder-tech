create table if not exists images
(
    id       bigserial primary key,
    filename text not null
);

insert into images(filename) values('file1'), ('file2'), ('file3');
