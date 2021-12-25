-- name: GetAccount :one
SELECT * FROM accounts 
WHERE id = ? LIMIT 1;

-- name: ListAccount :many
SELECT * FROM accounts 
ORDER BY owner;

-- name: CreateAccount :execresult
INSERT INTO accounts  (
    owner,
    balance,
    currency
) VALUES (
  ?, ?, ?
);

-- name: UpdateAccount :execresult
UPDATE accounts SET
    balance = ?, 
    currency = ?
WHERE id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE id = ?;