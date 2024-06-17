package db

import (
	"context"
	"testing"

	"github.com/Yelsnik/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createTestEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntries(t *testing.T) {
	acc := createTestAccount(t)

	createTestEntries(t, acc)
}

func TestGetEntries(t *testing.T) {
	account := createTestAccount(t)
	entry1 := createTestEntries(t, account)

	entry2, err := testQueries.GetEntries(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)

}

func TestListEntries(t *testing.T) {
	account := createTestAccount(t)

	for i := 0; i < 10; i++ {
		createTestEntries(t, account)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
