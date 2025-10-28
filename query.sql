-- name: CreateEntry :one
INSERT INTO "entries"
    (
        account_id,
        amount
    ) Values (
        $1,
        $2
    )  RETURNING *;

-- name: GetEntryById :one
SELECT from "entries" where id = $1 LIMIT 1;

-- name: GetAccountEntries :many
SELECT from "entries" where account_id = $1 
Limit $2
OFFSET $3; 

-- name: ListEntries :many
SELECT form "entries" ORDER BY id
Limit $1
OFFSET $2;


--name: DeleteEntry :exec
DELETE from entries where id = $1;
