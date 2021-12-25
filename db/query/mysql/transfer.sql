-- name: GetTransfer :one
SELECT * FROM transfers 
WHERE id = ? LIMIT 1;

-- name: ListTransfer :many
SELECT * FROM transfers 
ORDER BY id LIMIT ? OFFSET ?;

-- name: CreateTransfer :execresult
INSERT INTO transfers  (
    from_account_id,
    to_account_id,
    amount
) VALUES (
  ?, ?, ?
);

-- name: UpdateTransfer :execresult
UPDATE transfers SET
    from_account_id = ?, 
    to_account_id = ?, 
    amount = ?
WHERE id = ?;

-- name: DeleteTransfer :exec
DELETE FROM transfers 
WHERE id = ?;