// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: review.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (
  order_id,
  rating,
  review_date,
  created_at
) VALUES (
  $1, $2, $3, DEFAULT
) RETURNING review_id, order_id, created_at, rating, review_date
`

type CreateReviewParams struct {
	OrderID    int32       `json:"order_id"`
	Rating     int32       `json:"rating"`
	ReviewDate pgtype.Date `json:"review_date"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, createReview, arg.OrderID, arg.Rating, arg.ReviewDate)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.OrderID,
		&i.CreatedAt,
		&i.Rating,
		&i.ReviewDate,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE FROM reviews
WHERE review_id = $1
RETURNING review_id, order_id, created_at, rating, review_date
`

func (q *Queries) DeleteReview(ctx context.Context, reviewID int32) error {
	_, err := q.db.Exec(ctx, deleteReview, reviewID)
	return err
}

const getReview = `-- name: GetReview :one
SELECT review_id, order_id, created_at, rating, review_date FROM reviews
WHERE review_id = $1
`

func (q *Queries) GetReview(ctx context.Context, reviewID int32) (Review, error) {
	row := q.db.QueryRow(ctx, getReview, reviewID)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.OrderID,
		&i.CreatedAt,
		&i.Rating,
		&i.ReviewDate,
	)
	return i, err
}

const getReviewsByOrder = `-- name: GetReviewsByOrder :many
SELECT review_id, order_id, created_at, rating, review_date FROM reviews
WHERE order_id = $1
`

func (q *Queries) GetReviewsByOrder(ctx context.Context, orderID int32) ([]Review, error) {
	rows, err := q.db.Query(ctx, getReviewsByOrder, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Review{}
	for rows.Next() {
		var i Review
		if err := rows.Scan(
			&i.ReviewID,
			&i.OrderID,
			&i.CreatedAt,
			&i.Rating,
			&i.ReviewDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateReview = `-- name: UpdateReview :one
UPDATE reviews
SET 
  order_id = COALESCE($1, order_id),
  rating = COALESCE($2, rating),
  review_date = COALESCE($3, review_date)
WHERE 
  review_id = $4
RETURNING review_id, order_id, created_at, rating, review_date
`

type UpdateReviewParams struct {
	OrderID    pgtype.Int4 `json:"order_id"`
	Rating     pgtype.Int4 `json:"rating"`
	ReviewDate pgtype.Date `json:"review_date"`
	ReviewID   int32       `json:"review_id"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, updateReview,
		arg.OrderID,
		arg.Rating,
		arg.ReviewDate,
		arg.ReviewID,
	)
	var i Review
	err := row.Scan(
		&i.ReviewID,
		&i.OrderID,
		&i.CreatedAt,
		&i.Rating,
		&i.ReviewDate,
	)
	return i, err
}