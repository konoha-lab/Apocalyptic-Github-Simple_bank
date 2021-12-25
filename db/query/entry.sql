-- name: GetEntry :one
SELECT * FROM entries 
WHERE id = ? LIMIT 1;

-- name: ListEntry :many
SELECT * FROM entries 
ORDER BY owner LIMIT ? OFFSET ?;

-- name: CreateEntry :execresult
INSERT INTO entries  (
    account_id,
    amount
) VALUES (
  ?, ?
);

-- name: UpdateEntry :execresult
UPDATE entries SET
    amount = ?
WHERE id = ?;

-- name: DeleteEntry :exec
DELETE FROM entries 
WHERE id = ?;