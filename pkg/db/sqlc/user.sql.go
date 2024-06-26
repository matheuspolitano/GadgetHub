// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  email,
  hash_password,
  phone,
  user_role
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING user_id, first_name, last_name, email, created_at, hash_password, phone, user_role
`

type CreateUserParams struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
	Phone        string `json:"phone"`
	UserRole     string `json:"user_role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.HashPassword,
		arg.Phone,
		arg.UserRole,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.HashPassword,
		&i.Phone,
		&i.UserRole,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
where user_id = $1
RETURNING user_id, first_name, last_name, email, created_at, hash_password, phone, user_role
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUser, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, first_name, last_name, email, created_at, hash_password, phone, user_role FROM users
where user_id = $1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.HashPassword,
		&i.Phone,
		&i.UserRole,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, first_name, last_name, email, created_at, hash_password, phone, user_role FROM users
where email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.HashPassword,
		&i.Phone,
		&i.UserRole,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT user_id, first_name, last_name, email, created_at, hash_password, phone, user_role FROM users
where phone = $1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.HashPassword,
		&i.Phone,
		&i.UserRole,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
  hash_password = COALESCE($1,hash_password),
  first_name = COALESCE($2,first_name),
  last_name = COALESCE($3,last_name),
  email = COALESCE($4,email),
  phone = COALESCE($5,phone)
WHERE 
  user_id = $6
RETURNING user_id, first_name, last_name, email, created_at, hash_password, phone, user_role
`

type UpdateUserParams struct {
	HashPassword pgtype.Text `json:"hash_password"`
	FirstName    pgtype.Text `json:"first_name"`
	LastName     pgtype.Text `json:"last_name"`
	Email        pgtype.Text `json:"email"`
	Phone        pgtype.Text `json:"phone"`
	UserID       int32       `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.HashPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.UserID,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.HashPassword,
		&i.Phone,
		&i.UserRole,
	)
	return i, err
}
