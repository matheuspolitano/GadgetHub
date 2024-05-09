-- name: CreateOrder :one
INSERT INTO orders (
  product_id,
  user_id,
  coupon_id,
  price,
  created_at
) VALUES (
  $1, $2, $3, $4, DEFAULT
) RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE order_id = $1;

-- name: GetOrdersByUser :many
SELECT * FROM orders
WHERE user_id = $1;

-- name: GetOrdersByProduct :many
SELECT * FROM orders
WHERE product_id = $1;

-- name: UpdateOrder :one
UPDATE orders
SET 
  product_id = COALESCE(sqlc.narg(product_id), product_id),
  user_id = COALESCE(sqlc.narg(user_id), user_id),
  coupon_id = COALESCE(sqlc.narg(coupon_id), coupon_id),
  price = COALESCE(sqlc.narg(price), price)
WHERE 
  order_id = sqlc.arg(order_id)
RETURNING *;
