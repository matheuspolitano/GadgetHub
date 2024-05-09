-- name: CreateDiscountCoupon :one
INSERT INTO discount_coupons (
  created_by,
  discount,
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