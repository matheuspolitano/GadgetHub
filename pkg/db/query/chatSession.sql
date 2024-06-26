-- name: CreateChatSession :one
INSERT INTO chat_sessions (
  last_message_id,
  action_flow,
  user_id,
  payload,
  opened_at,
  closed_at
) VALUES (
  $1, $2, $3, $4, DEFAULT, $5
) RETURNING *;

-- name: DeleteChatSession :exec
DELETE FROM chat_sessions
WHERE chat_session_id = $1
RETURNING *;

-- name: GetChatSession :one
SELECT * FROM chat_sessions
WHERE chat_session_id = $1;

-- name: GetChatSessionsByUser :many
SELECT * FROM chat_sessions
WHERE user_id = $1;

-- name: GetChatSessionsByUserPhone :one
SELECT 
    cs.chat_session_id,
    cs.last_message_id,
    cs.action_flow,
    cs.user_id,
    cs.payload,
    cs.opened_at,
    cs.closed_at
FROM chat_sessions cs
JOIN users u ON cs.user_id = u.user_id
WHERE u.phone = $1
  AND cs.closed_at IS NULL
ORDER BY cs.opened_at DESC
LIMIT 1;


-- name: UpdateChatSession :one
UPDATE chat_sessions
SET 
  last_message_id = COALESCE(sqlc.narg(last_message_id), last_message_id),
  action_flow = COALESCE(sqlc.narg(action_flow), action_flow),
  user_id = COALESCE(sqlc.narg(user_id), user_id),
  payload = COALESCE(sqlc.narg(payload), payload),
  opened_at = COALESCE(sqlc.narg(opened_at), opened_at),
  closed_at = COALESCE(sqlc.narg(closed_at), closed_at)
WHERE 
  chat_session_id = sqlc.arg(chat_session_id)
RETURNING *;