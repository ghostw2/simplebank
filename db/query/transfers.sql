-- name: CreateTransfer :one
INSERT into "transfers"  (
    from_account_id,
    to_account_id,
    amount)
    Values (
        $1,
        $2,
        $3
    ) RETURNING *;

-- name: ListTransfers :many
SELECT * from "transfers" ORDER BY id
    LIMIT $1
    OFFSET $2;
-- name: GetTranfer :one
SELECT * from "transfers" WHERE id = $1 LIMIT 1;

-- name: DeleteTransfer :exec
DELETE from "transfers" where id = $1;

