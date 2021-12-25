-- name: GetAccount :one
SELECT * FROM accounts 
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts 
ORDER BY owner LIMIT ? OFFSET ?;

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