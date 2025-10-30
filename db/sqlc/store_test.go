package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t, testQueries, context.Background())
	account2 := CreateRandomAccount(t, testQueries, context.Background())
	assert.NotNil(t, testStore)
	result, err := testStore.Transfer(context.Background(),
		TransferTxParam{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        200,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, result)

}
