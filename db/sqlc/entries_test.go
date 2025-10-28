package db

import (
	"context"
	"fmt"
	utilsdb "menribardhi/micro-go-psql/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRandomEntry(t *testing.T, args CreateEntryParams) Entry {
	createdEntry, err := testQueries.CreateEntry(context.Background(), args)
	assert.Nil(t, err)
	assert.NotNil(t, createdEntry)
	assert.Equal(t, args.AccountID, createdEntry.AccountID)
	assert.Equal(t, args.Amount, createdEntry.Amount)
	return createdEntry
}
func TestCreateEntry(t *testing.T) {
	Account := createRandomAccount(t, testQueries, context.Background())
	createRandomEntry(t, CreateEntryParams{
		AccountID: Account.ID,
		Amount:    int64(utilsdb.RandomInt(0, 900)),
	})
}

func retriveEntries(t *testing.T) []ListEntriesRow {
	entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{Limit: 99999, Offset: 0})
	assert.Nil(t, err)
	return entries
}

func TestCreateManyEntries(t *testing.T) {
	intialEntries := retriveEntries(t)
	var count = 3
	account := createRandomAccount(t, testQueries, context.Background())
	for range count {
		testQueries.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: account.ID,
			Amount:    int64(utilsdb.RandomInt(0, 900)),
		})
	}

	finalEntries := retriveEntries(t)

	assert.Equal(t, len(intialEntries)+count, len(finalEntries), fmt.Sprintf("inital entries %v ,final entries %v", len(intialEntries), len(finalEntries)))

}

func deleteEntry_test(t *testing.T, id int64) {
	err := testQueries.DeleteEntry(context.Background(), id)
	assert.Nil(t, err)
}

func TestEntryDeletion(t *testing.T) {
	var count = 3
	account := createRandomAccount(t, testQueries, context.Background())

	intialEntries := retriveEntries(t)
	var entries []Entry
	var entry Entry
	var err error

	for range count {
		entry, err = testQueries.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: account.ID,
			Amount:    int64(utilsdb.RandomInt(0, 900)),
		})
		assert.Nil(t, err)
		entries = append(entries, entry)
	}
	for _, entry := range entries {
		deleteEntry_test(t, entry.ID)
	}
	finalEntries := retriveEntries(t)
	assert.Equal(t, len(intialEntries), len(finalEntries), fmt.Sprintf("inital entries %v ,final entries %v", len(intialEntries), len(finalEntries)))

	for _, entry = range entries {
		entry, err := testQueries.GetEntryById(context.Background(), entry.ID)
		assert.NotNil(t, err)
		assert.Equal(t, entry.ID, int64(0))
	}
}
