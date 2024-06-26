-- name: CreateUser :one
insert into users (email, password, name, gender, age)
values ($1, $2, $3, $4, $5)
returning id, email, password, name, gender, age;

-- name: GetUser :one
select *
from users
where email = $1
limit 1;
