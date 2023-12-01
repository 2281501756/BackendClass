package db

import (
	"context"
	"database/sql"
	"errors"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}

func (store *Store) TransferMoney(account1, account2, amount int32) error {
	return store.execTx(context.Background(), func(queries *Queries) error {
		user1, err := queries.GetUserForUpdate(context.Background(), account1)
		if err != nil {
			return err
		}
		user2, err := queries.GetUserForUpdate(context.Background(), account2)
		if user1.Money < amount {
			return errors.New("没钱")
		}
		user1, err = queries.UpdateUserMoney(context.Background(), UpdateUserMoneyParams{
			Money: -amount,
			ID:    user1.ID,
		})
		if err != nil {
			return err
		}
		user2, err = queries.UpdateUserMoney(context.Background(), UpdateUserMoneyParams{
			Money: amount,
			ID:    user2.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
