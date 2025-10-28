package db

import (
	"context"
	utilsdb "menribardhi/micro-go-psql/utils"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRandomAccount(t *testing.T, queries *Queries, ctx context.Context) Account {
	args := CreateAccountParams{
		Owner:    utilsdb.RandomString(5),
		Balance:  utilsdb.RandomString(3),
		Currency: utilsdb.RandomString(4),
	}
	account, err := queries.CreateAccount(ctx, args)
	assert.Nil(t, err)
	assert.Equal(t, args.Owner, account.Owner)
	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t, testQueries, context.Background())
}
func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t, testQueries, context.Background())
	updateAccountWithId(t, createdAccount.ID)
}
func updateAccountWithId(t *testing.T, id int64) {
	args := UpdateAccountParams{
		ID:      id,
		Balance: strconv.Itoa(utilsdb.RandomInt(0, 30000)),
	}
	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)
	assert.Nil(t, err)
	assert.NotNil(t, updatedAccount)
	assert.Equal(t, args.Balance, updatedAccount.Balance)
}
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t, testQueries, context.Background())
	deleteAccountbyId(t, account.ID)
}

func deleteAccountbyId(t *testing.T, id int64) {
	existingAccount, err := testQueries.GetAccount(context.Background(), id)
	assert.Nil(t, err)
	assert.NotNil(t, existingAccount)

	deleteErr := testQueries.DeleteAccount(context.Background(), existingAccount.ID)

	assert.Nil(t, deleteErr)

}
