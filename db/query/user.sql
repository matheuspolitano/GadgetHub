-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email,
  hash_password,
  phone,
  is_admin
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
where user_id = $1;