

create table buckets (
    id serial primary key,
    name varchar(100) not null unique,
    created timestamp not null,
    updated timestamp not null
);

create table files (
    id serial primary key,
    name text not null,
    bucket_id integer not null references buckets(id) on delete cascade,
    size integer not null,
    sizeAfterCompression integer not null,
    extension varchar(10) not null,
    mimeType text not null,
    created timestamp not null,
    updated timestamp not null,
    files_token TSVECTOR,
);