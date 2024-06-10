-- users
create table users (
    id serial primary key,
    email text not null unique,
    password text not null,
    name text not null,
    gender text not null,
    age int not null
);