package db

import (
	"context"
	"database/sql"
	"fmt"
)

//store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

//store provides all functions to execute SQL queries and transactions
type SqlStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SqlStore) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountNumber int64 `json:"from_account_number"`
	ToAccountNumber   int64 `json:"to_account_number"`
	Amount            int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *SqlStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAccountNumber: arg.FromAccountNumber,
			ToAccountNumber:   arg.ToAccountNumber,
			Amount:            arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountNumber: arg.FromAccountNumber,
			Amount:        -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountNumber: arg.ToAccountNumber,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountNumber > arg.ToAccountNumber {
			result.FromAccount, result.ToAccount, _ = addMoney(ctx, q, arg.FromAccountNumber, -arg.Amount, arg.ToAccountNumber, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, _ = addMoney(ctx, q, arg.ToAccountNumber, arg.Amount, arg.FromAccountNumber, -arg.Amount)

		}

		return nil
	})
	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	account1AccountNumber int64,
	amount1 int64,
	account2AccountNumber int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		AccountNumber: account1AccountNumber,
		Amount:        amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		AccountNumber: account2AccountNumber,
		Amount:        amount2,
	})
	if err != nil {
		return
	}

	return
}
