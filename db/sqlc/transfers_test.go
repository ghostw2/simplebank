package db

import (
	"context"
	"fmt"
	utilsdb "menribardhi/micro-go-psql/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func deleteTransfer_test(t *testing.T, id int64) {
	err := testQueries.DeleteTransfer(context.Background(), id)
	assert.Nil(t, err)
}
func createTransfer_test(t *testing.T, args CreateTransferParams) Transfer {
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	assert.Nil(t, err)
	assert.NotNil(t, transfer)
	return transfer
}
func retriveTransfers_test(t *testing.T) []Transfer {
	transfers, err := testQueries.ListTransfers(context.Background(), ListTransfersParams{Limit: 99999, Offset: 0})
	assert.Nil(t, err)
	assert.NotEmpty(t, transfers)
	return transfers
}

func TestCreateTransfer(t *testing.T) {

	account1 := CreateRandomAccount(t, testQueries, context.Background())
	account2 := CreateRandomAccount(t, testQueries, context.Background())
	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        int64(utilsdb.RandomInt(0, 3000)),
	}
	transfer := createTransfer_test(t, args)
	assert.Equal(t, transfer.Amount, args.Amount)
	assert.Equal(t, transfer.ToAccountID, args.ToAccountID)
	assert.Equal(t, transfer.FromAccountID, args.FromAccountID)
}
func TestCreateManyTransfers(t *testing.T) {
	intial := retriveTransfers_test(t)
	var count = 3
	account1 := CreateRandomAccount(t, testQueries, context.Background())
	account2 := CreateRandomAccount(t, testQueries, context.Background())
	for range count {
		args := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        int64(utilsdb.RandomInt(0, 3000)),
		}
		createTransfer_test(t, args)
	}

	final := retriveTransfers_test(t)

	assert.Equal(t, len(intial)+count, len(final), fmt.Sprintf("inital entries %v ,final entries %v", len(intial), len(final)))

}

func TestTransferDeletion(t *testing.T) {
	intial := retriveTransfers_test(t)
	var count = 3

	var transfers []Transfer
	var transfer Transfer
	var err error

	account1 := CreateRandomAccount(t, testQueries, context.Background())
	account2 := CreateRandomAccount(t, testQueries, context.Background())
	for range count {
		args := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        int64(utilsdb.RandomInt(0, 3000)),
		}
		transfer = createTransfer_test(t, args)
		transfers = append(transfers, transfer)
		assert.Nil(t, err)
	}
	for _, tr := range transfers {
		deleteTransfer_test(t, tr.ID)
	}
	final := retriveTransfers_test(t)
	assert.Equal(t, len(intial), len(final), fmt.Sprintf("inital entries %v ,final entries %v", len(intial), len(final)))

	for _, tr := range transfers {
		entry, err := testQueries.GetTranfer(context.Background(), tr.ID)
		assert.NotNil(t, err)
		assert.Equal(t, entry.ID, int64(0))
	}

}
