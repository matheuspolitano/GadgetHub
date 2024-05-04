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
where user_id = $1
RETURNING *;


-- name: GetUser :one
SELECT * FROM users
where user_id = $1;


-- name: GetUserByEmail :one
SELECT * FROM users
where email = $1;

-- name: GetUserByPhone :one
SELECT * FROM users
where phone = $1;

-- name: UpdateUser :one
UPDATE users
SET 
  hash_password = COALESCE(sqlc.narg(hash_password),hash_password),
  first_name = COALESCE(sqlc.narg(first_name),first_name),
  last_name = COALESCE(sqlc.narg(last_name),last_name),
  email = COALESCE(sqlc.narg(email),email),
  phone = COALESCE(sqlc.narg(phone),phone)
WHERE 
  user_id = sqlc.arg(user_id)
RETURNING *;  
