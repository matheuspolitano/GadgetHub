-- name: CreateChatMessage :one
INSERT INTO chat_messages (
  chat_session_id,
  message_received,
  message_sent,
  received_at,
  sent_at,
  action,
  message_before_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: DeleteChatMessage :exec
DELETE FROM chat_messages
WHERE chat_message_id = $1
RETURNING *;

-- name: GetChatMessage :one
SELECT * FROM chat_messages
WHERE chat_message_id = $1;

-- name: GetChatMessagesBySession :many
SELECT * FROM chat_messages
WHERE chat_session_id = $1;

-- name: UpdateChatMessage :one
UPDATE chat_messages
SET 
  chat_session_id = COALESCE(sqlc.narg(chat_session_id), chat_session_id),
  message_received = COALESCE(sqlc.narg(message_received), message_received),
  message_sent = COALESCE(sqlc.narg(message_sent), message_sent),
  received_at = COALESCE(sqlc.narg(received_at), received_at),
  sent_at = COALESCE(sqlc.narg(sent_at), sent_at),
  action = COALESCE(sqlc.narg(action), action),
  message_before_id = COALESCE(sqlc.narg(message_before_id), message_before_id)
WHERE 
  chat_message_id = sqlc.arg(chat_message_id)
RETURNING *;
