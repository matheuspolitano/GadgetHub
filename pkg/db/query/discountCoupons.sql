-- name: CreateDiscountCoupon :one
INSERT INTO discount_coupons (
  created_by,
  created_at,
  expires_at
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteDiscountCoupon :exec
DELETE FROM discount_coupons
WHERE coupon_id = $1
RETURNING *;

-- name: GetDiscountCoupon :one
SELECT * FROM discount_coupons
WHERE coupon_id = $1;

-- name: GetCouponsByCreator :many
SELECT * FROM discount_coupons
WHERE created_by = $1;

-- name: UpdateDiscountCoupon :one
UPDATE discount_coupons
SET 
  created_by = COALESCE(sqlc.narg(created_by), created_by),
  created_at = COALESCE(sqlc.narg(created_at), created_at),
  expires_at = COALESCE(sqlc.narg(expires_at), expires_at)
WHERE 
  coupon_id = sqlc.arg(coupon_id)
RETURNING *;
