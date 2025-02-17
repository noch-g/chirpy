-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES ($1, now(), now(), $2, $3, null)
RETURNING *;

-- name: GetRefreshToken :one
SELECT *
FROM refresh_tokens
WHERE token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET updated_at = now(), revoked_at = now()
WHERE token = $1;
