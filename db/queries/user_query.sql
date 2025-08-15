-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateUserBalance :exec
UPDATE users
SET balance = $1, updated_at = NOW()
WHERE id = $2
AND deleted_at IS NULL;