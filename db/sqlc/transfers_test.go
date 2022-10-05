package db

import (
	"context"
	"testing"
	"time"

	util "github.com/mayowa1305/simpleBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTranfers(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransfersParams{
		FromAccountNumber: account1.AccountNumber,
		ToAccountNumber:   account2.AccountNumber,
		Amount:            util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountNumber, transfer.FromAccountNumber)
	require.Equal(t, arg.ToAccountNumber, transfer.ToAccountNumber)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTranfers(t, account1, account2)
}

func TestGeTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTranfers(t, account1, account2)
	transfer2, err := testQueries.GetTransfers(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountNumber, transfer2.FromAccountNumber)
	require.Equal(t, transfer1.ToAccountNumber, transfer2.ToAccountNumber)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTranfers(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountNumber: account1.AccountNumber,
		ToAccountNumber:   account2.AccountNumber,
		Limit:             5,
		Offset:            0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountNumber == account1.AccountNumber || transfer.ToAccountNumber == account1.AccountNumber)
	}
}
