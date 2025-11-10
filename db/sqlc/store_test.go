package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t, testQueries, context.Background())
	account2 := CreateRandomAccount(t, testQueries, context.Background())
	assert.NotNil(t, testStore)

	fmt.Println(">> Balance before ;", account1.Balance, account2.Balance)

	n := 10
	amount := 100
	errs := make(chan error)
	for i := 0; i < n; i++ {
		fromAccountid := account1.ID
		toAccontId := account2.ID
		if i%2 == 1 {
			fromAccountid = account2.ID
			toAccontId = account1.ID
		}
		go func() {
			_, err := testStore.Transfer(context.Background(),
				TransferTxParam{
					FromAccountID: fromAccountid,
					ToAccountID:   toAccontId,
					Amount:        int64(amount),
				},
			)
			errs <- err
		}()

	}
	for i := 0; i < n; i++ {
		err := <-errs
		assert.Nil(t, err)
		// assert.NotNil(t, result)
	}
	account1, _ = testQueries.GetAccount(context.Background(), account1.ID)
	account2, _ = testQueries.GetAccount(context.Background(), account2.ID)
	fmt.Println(">> Balance after ;", account1.Balance, account2.Balance)

}
