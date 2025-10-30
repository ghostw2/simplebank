package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error :%v rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

type TransferTxParam struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) anotherExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rerr := tx.Rollback(ctx); rerr != nil {
			return fmt.Errorf("transaction error :%v and the rollback error:%v", err, rerr)
		}
		return err
	}
	return tx.Commit(ctx)
}

func (store *Store) Transfer(ctx context.Context, args TransferTxParam) (TransferTxResult, error) {
	//deduct from first account
	var result TransferTxResult
	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error
		fromAccount, err := q.GetAccount(ctx, args.FromAccountID)
		if err != nil {
			return err
		}
		result.FromAccount, err = q.UpdateAccount(ctx,
			UpdateAccountParams{
				ID:      args.FromAccountID,
				Balance: fromAccount.Balance - args.Amount,
			})
		if err != nil {
			return err
		}
		//add from secodnd account
		toAccount, err := q.GetAccount(ctx, args.ToAccountID)
		if err != nil {
			return err
		}
		result.ToAccount, err = q.UpdateAccount(ctx,
			UpdateAccountParams{
				ID:      args.ToAccountID,
				Balance: toAccount.Balance + args.Amount,
			})
		if err != nil {
			return err
		}
		//entry 1
		result.FromEntry, err = q.CreateEntry(ctx,
			CreateEntryParams{
				AccountID: args.FromAccountID,
				Amount:    -args.Amount,
			})
		if err != nil {
			return err
		}
		//entry 2
		result.ToEntry, err = q.CreateEntry(ctx,
			CreateEntryParams{
				AccountID: args.ToAccountID,
				Amount:    args.Amount,
			})
		if err != nil {
			return err
		}
		//transfer
		result.Transfer, err = q.CreateTransfer(ctx,
			CreateTransferParams{
				FromAccountID: args.FromAccountID,
				ToAccountID:   args.ToAccountID,
				Amount:        args.Amount,
			})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
