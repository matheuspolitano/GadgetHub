-- name: CreateReview :one
INSERT INTO reviews (
  order_id,
  rating,
  review_date,
  created_at
) VALUES (
  $1, $2, $3, DEFAULT
) RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE review_id = $1
RETURNING *;

-- name: GetReview :one
SELECT * FROM reviews
WHERE review_id = $1;

-- name: GetReviewsByOrder :many
SELECT * FROM reviews
WHERE order_id = $1;
