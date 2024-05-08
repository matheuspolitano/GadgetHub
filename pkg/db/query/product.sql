-- name: CreateProduct :one
INSERT INTO products (
  name,
  description,
  price,
  stock,
  category_id,
  brand,
  model
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE product_id = $1;

-- name: GetProductsByCategory :many
SELECT * FROM products
WHERE category_id = $1;

-- name: UpdateProduct :one
UPDATE products
SET 
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description),
  price = COALESCE(sqlc.narg(price), price),
  stock = COALESCE(sqlc.narg(stock), stock),
  category_id = COALESCE(sqlc.narg(category_id), category_id),
  brand = COALESCE(sqlc.narg(brand), brand),
  model = COALESCE(sqlc.narg(model), model)
WHERE 
  product_id = sqlc.arg(product_id)
RETURNING *;
