// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: transfers.sql

package db

import (
	"context"
)

const createTransfers = `-- name: CreateTransfers :one
INSERT INTO "transfers" (
  "from_account_id",
  "to_account_id",
  "amount"
) VALUES (
  $1, $2, $3
)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransfersParams struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfer, error) {
	row := q.queryRow(ctx, q.createTransfersStmt, createTransfers, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfers = `-- name: GetTransfers :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM "transfers"
WHERE "id" = $1 LIMIT 1
`

func (q *Queries) GetTransfers(ctx context.Context, id int64) (Transfer, error) {
	row := q.queryRow(ctx, q.getTransfersStmt, getTransfers, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM "transfers"
WHERE
 "from_account_id" = $1 or
 "to_account_id" = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfersParams struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.query(ctx, q.listTransfersStmt, listTransfers,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
