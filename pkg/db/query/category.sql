-- name: CreateCategory :one
INSERT INTO categories (
  name,
  description
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE category_id = $1
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE category_id = $1;

-- name: GetCategoriesByName :many
SELECT * FROM categories
WHERE name ILIKE $1;  -- Allows searching by name with case-insensitive matching.

-- name: UpdateCategory :one
UPDATE categories
SET 
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description)
WHERE 
  category_id = sqlc.arg(category_id)
RETURNING *;
