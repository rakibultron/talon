-- name: GetAllUsers :many
SELECT id, name, email, password, created_at
FROM users;


-- name: GetUserByID :one
SELECT id, name, email, password, created_at
FROM users
WHERE id = $1;


-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password, created_at;


-- name: UpdateUserEmail :exec
UPDATE users
SET email = $2
WHERE id = $1;
