-- name: CreateAccount :one
INSERT INTO "accounts" (
    owner,
    balance,
    currency
) Values (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetAccount :one
SELECT * from "accounts" where id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * from "accounts" ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts SET balance = $2 
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE from accounts WHERE id = $1;

-- name: GetAccountFormUpdate :one
SELECT * from "accounts" where id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: AddBalance :exec
UPDATE accounts set balance = balance + $1 where id = $2;