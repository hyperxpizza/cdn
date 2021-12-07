
create table files (
    id serial primary key,
    name text not null,
    path text not null,
    size integer not null,
    sizeAfterCompression integer not null,
    extension varchar(10) not null,
    created timestamp not null,
    updated timestamp not null
);