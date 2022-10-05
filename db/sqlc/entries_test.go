package db

import (
	"context"
	"testing"
	"time"

	util "github.com/mayowa1305/simpleBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountNumber: account.AccountNumber,
		Amount:        util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountNumber, entry.AccountNumber)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntries(t, account)
}

func TestGetEntries(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntries(t, account)
	entry2, err := testQueries.GetEntries(context.Background(), entry1.AccountNumber)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountNumber, entry2.AccountNumber)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	Account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntries(t, Account)
	}

	arg := ListEntriesParams{
		AccountNumber: Account.AccountNumber,
		Limit:         5,
		Offset:        0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entries := range entries {
		require.NotEmpty(t, entries)
		require.Equal(t, arg.AccountNumber, entries.AccountNumber)
	}
}
