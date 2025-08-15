-- name: CreateTransaction :exec
INSERT INTO transactions (
        user_id,
        amount,
        type,
        status,
        description,
        transaction_date
    )
VALUES ($1, $2, $3, $4, $5, NOW());

-- name: GetUserTransactions :many
SELECT *
FROM transactions
WHERE user_id = $1
    and deleted_at IS NULL
ORDER BY transaction_date DESC;