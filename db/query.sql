-- name: GetAccounts :many
SELECT * FROM accounts;

-- name: UpdateAccount :exec
UPDATE accounts
SET
        name = $1,
        notes = $2,
        align = $3
WHERE id = $4;
