-- name: CreateEntries :one
INSERT INTO "entries" (
  "account_number",
  "amount"
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntries :one
SELECT * FROM "entries"
WHERE "account_number" = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM "entries"
WHERE "account_number" = $1
ORDER BY "id"
LIMIT $2
OFFSET $3;