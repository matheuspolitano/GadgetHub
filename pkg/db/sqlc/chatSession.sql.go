// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: chatSession.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createChatSession = `-- name: CreateChatSession :one
INSERT INTO chat_sessions (
  last_message_id,
  action_flow,
  user_id,
  payload,
  opened_at,
  closed_at
) VALUES (
  $1, $2, $3, $4, DEFAULT, $5
) RETURNING chat_session_id, last_message_id, action_flow, user_id, payload, opened_at, closed_at
`

type CreateChatSessionParams struct {
	LastMessageID pgtype.Int4        `json:"last_message_id"`
	ActionFlow    string             `json:"action_flow"`
	UserID        int32              `json:"user_id"`
	Payload       string             `json:"payload"`
	ClosedAt      pgtype.Timestamptz `json:"closed_at"`
}

func (q *Queries) CreateChatSession(ctx context.Context, arg CreateChatSessionParams) (ChatSession, error) {
	row := q.db.QueryRow(ctx, createChatSession,
		arg.LastMessageID,
		arg.ActionFlow,
		arg.UserID,
		arg.Payload,
		arg.ClosedAt,
	)
	var i ChatSession
	err := row.Scan(
		&i.ChatSessionID,
		&i.LastMessageID,
		&i.ActionFlow,
		&i.UserID,
		&i.Payload,
		&i.OpenedAt,
		&i.ClosedAt,
	)
	return i, err
}

const deleteChatSession = `-- name: DeleteChatSession :exec
DELETE FROM chat_sessions
WHERE chat_session_id = $1
RETURNING chat_session_id, last_message_id, action_flow, user_id, payload, opened_at, closed_at
`

func (q *Queries) DeleteChatSession(ctx context.Context, chatSessionID int32) error {
	_, err := q.db.Exec(ctx, deleteChatSession, chatSessionID)
	return err
}

const getChatSession = `-- name: GetChatSession :one
SELECT chat_session_id, last_message_id, action_flow, user_id, payload, opened_at, closed_at FROM chat_sessions
WHERE chat_session_id = $1
`

func (q *Queries) GetChatSession(ctx context.Context, chatSessionID int32) (ChatSession, error) {
	row := q.db.QueryRow(ctx, getChatSession, chatSessionID)
	var i ChatSession
	err := row.Scan(
		&i.ChatSessionID,
		&i.LastMessageID,
		&i.ActionFlow,
		&i.UserID,
		&i.Payload,
		&i.OpenedAt,
		&i.ClosedAt,
	)
	return i, err
}

const getChatSessionsByUser = `-- name: GetChatSessionsByUser :many
SELECT chat_session_id, last_message_id, action_flow, user_id, payload, opened_at, closed_at FROM chat_sessions
WHERE user_id = $1
`

func (q *Queries) GetChatSessionsByUser(ctx context.Context, userID int32) ([]ChatSession, error) {
	rows, err := q.db.Query(ctx, getChatSessionsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ChatSession{}
	for rows.Next() {
		var i ChatSession
		if err := rows.Scan(
			&i.ChatSessionID,
			&i.LastMessageID,
			&i.ActionFlow,
			&i.UserID,
			&i.Payload,
			&i.OpenedAt,
			&i.ClosedAt,
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

const updateChatSession = `-- name: UpdateChatSession :one
UPDATE chat_sessions
SET 
  last_message_id = COALESCE($1, last_message_id),
  action_flow = COALESCE($2, action_flow),
  user_id = COALESCE($3, user_id),
  payload = COALESCE($4, payload),
  opened_at = COALESCE($5, opened_at),
  closed_at = COALESCE($6, closed_at)
WHERE 
  chat_session_id = $7
RETURNING chat_session_id, last_message_id, action_flow, user_id, payload, opened_at, closed_at
`

type UpdateChatSessionParams struct {
	LastMessageID pgtype.Int4        `json:"last_message_id"`
	ActionFlow    pgtype.Text        `json:"action_flow"`
	UserID        pgtype.Int4        `json:"user_id"`
	Payload       pgtype.Text        `json:"payload"`
	OpenedAt      pgtype.Timestamptz `json:"opened_at"`
	ClosedAt      pgtype.Timestamptz `json:"closed_at"`
	ChatSessionID int32              `json:"chat_session_id"`
}

func (q *Queries) UpdateChatSession(ctx context.Context, arg UpdateChatSessionParams) (ChatSession, error) {
	row := q.db.QueryRow(ctx, updateChatSession,
		arg.LastMessageID,
		arg.ActionFlow,
		arg.UserID,
		arg.Payload,
		arg.OpenedAt,
		arg.ClosedAt,
		arg.ChatSessionID,
	)
	var i ChatSession
	err := row.Scan(
		&i.ChatSessionID,
		&i.LastMessageID,
		&i.ActionFlow,
		&i.UserID,
		&i.Payload,
		&i.OpenedAt,
		&i.ClosedAt,
	)
	return i, err
}
