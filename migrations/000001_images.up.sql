create table if not exists images
(
    id       bigserial primary key,
    filename text not null
);

INSERT INTO images(filename)
VALUES
    ('d84cef4b-adf5-401d-a945-d746c99cd43e'),
    ('c284f294-0635-4e98-a7be-110c3015afb5'),
    ('0203cb04-bc81-4775-ba77-8153385122af');
