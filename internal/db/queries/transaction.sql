-- name: CreateTransaction :one
INSERT INTO transactions (user_id, amount, type, status)
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: UpdateTransaction :one
UPDATE transactions
SET status = $2
WHERE id = $1
RETURNING *;
