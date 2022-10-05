-- name: CreateAccounts :one
INSERT INTO "accounts" (
  "owner",
  "balance",
  "currency"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM "accounts"
WHERE "account_number" = $1 LIMIT 1;

-- name: GetAccountforupdate :one
SELECT * FROM "accounts"
WHERE "account_number" = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM "accounts"
WHERE "owner" = $1
ORDER BY "account_number"
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE "accounts"
SET "balance" = $2
WHERE "account_number" = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE "accounts"
SET "balance" = "balance" + sqlc.arg(amount)
WHERE "account_number" = sqlc.arg(account_number)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM "accounts"
WHERE "account_number" = $1;