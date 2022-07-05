// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addAccountBalanceStmt, err = db.PrepareContext(ctx, addAccountBalance); err != nil {
		return nil, fmt.Errorf("error preparing query AddAccountBalance: %w", err)
	}
	if q.createAccountsStmt, err = db.PrepareContext(ctx, createAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccounts: %w", err)
	}
	if q.createEntriesStmt, err = db.PrepareContext(ctx, createEntries); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEntries: %w", err)
	}
	if q.createTransfersStmt, err = db.PrepareContext(ctx, createTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransfers: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getAccountforupdateStmt, err = db.PrepareContext(ctx, getAccountforupdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountforupdate: %w", err)
	}
	if q.getEntriesStmt, err = db.PrepareContext(ctx, getEntries); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntries: %w", err)
	}
	if q.getTransfersStmt, err = db.PrepareContext(ctx, getTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransfers: %w", err)
	}
	if q.listAccountsStmt, err = db.PrepareContext(ctx, listAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccounts: %w", err)
	}
	if q.listEntriesStmt, err = db.PrepareContext(ctx, listEntries); err != nil {
		return nil, fmt.Errorf("error preparing query ListEntries: %w", err)
	}
	if q.listTransfersStmt, err = db.PrepareContext(ctx, listTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query ListTransfers: %w", err)
	}
	if q.updateAccountStmt, err = db.PrepareContext(ctx, updateAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccount: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAccountBalanceStmt != nil {
		if cerr := q.addAccountBalanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAccountBalanceStmt: %w", cerr)
		}
	}
	if q.createAccountsStmt != nil {
		if cerr := q.createAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountsStmt: %w", cerr)
		}
	}
	if q.createEntriesStmt != nil {
		if cerr := q.createEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEntriesStmt: %w", cerr)
		}
	}
	if q.createTransfersStmt != nil {
		if cerr := q.createTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransfersStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getAccountforupdateStmt != nil {
		if cerr := q.getAccountforupdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountforupdateStmt: %w", cerr)
		}
	}
	if q.getEntriesStmt != nil {
		if cerr := q.getEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntriesStmt: %w", cerr)
		}
	}
	if q.getTransfersStmt != nil {
		if cerr := q.getTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransfersStmt: %w", cerr)
		}
	}
	if q.listAccountsStmt != nil {
		if cerr := q.listAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccountsStmt: %w", cerr)
		}
	}
	if q.listEntriesStmt != nil {
		if cerr := q.listEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listEntriesStmt: %w", cerr)
		}
	}
	if q.listTransfersStmt != nil {
		if cerr := q.listTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTransfersStmt: %w", cerr)
		}
	}
	if q.updateAccountStmt != nil {
		if cerr := q.updateAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                      DBTX
	tx                      *sql.Tx
	addAccountBalanceStmt   *sql.Stmt
	createAccountsStmt      *sql.Stmt
	createEntriesStmt       *sql.Stmt
	createTransfersStmt     *sql.Stmt
	deleteAccountStmt       *sql.Stmt
	getAccountStmt          *sql.Stmt
	getAccountforupdateStmt *sql.Stmt
	getEntriesStmt          *sql.Stmt
	getTransfersStmt        *sql.Stmt
	listAccountsStmt        *sql.Stmt
	listEntriesStmt         *sql.Stmt
	listTransfersStmt       *sql.Stmt
	updateAccountStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                      tx,
		tx:                      tx,
		addAccountBalanceStmt:   q.addAccountBalanceStmt,
		createAccountsStmt:      q.createAccountsStmt,
		createEntriesStmt:       q.createEntriesStmt,
		createTransfersStmt:     q.createTransfersStmt,
		deleteAccountStmt:       q.deleteAccountStmt,
		getAccountStmt:          q.getAccountStmt,
		getAccountforupdateStmt: q.getAccountforupdateStmt,
		getEntriesStmt:          q.getEntriesStmt,
		getTransfersStmt:        q.getTransfersStmt,
		listAccountsStmt:        q.listAccountsStmt,
		listEntriesStmt:         q.listEntriesStmt,
		listTransfersStmt:       q.listTransfersStmt,
		updateAccountStmt:       q.updateAccountStmt,
	}
}