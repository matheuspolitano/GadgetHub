-- name: CreateUser :one
INSERT INTO users (
  fist_name,
  last_name,
  email,
  hash_password,
  phone,
  is_admin
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;
