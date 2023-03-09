package main

import (
	"context"
	"io"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type GameUserOperation interface {
	createUser(context.Context, io.Writer, userParams) error
	addItemToUser(context.Context, io.Writer, userParams, itemParams) error
	userItems(context.Context, io.Writer, string) (*spanner.ReadOnlyTransaction, *spanner.RowIterator, error)
	listItems(context.Context) ([]itemParams, error)
}

type userParams struct {
	userID   string
	userName string
}

type itemParams struct {
	itemID string
}

type dbClient struct {
	sc *spanner.Client
}

func newClient(ctx context.Context, dbString string) (dbClient, error) {

	client, err := spanner.NewClient(ctx, dbString)
	if err != nil {
		return dbClient{}, err
	}
	return dbClient{
		sc: client,
	}, nil
}

// create a user
func (d dbClient) createUser(ctx context.Context, w io.Writer, u userParams) error {

	_, err := d.sc.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		sqlToUsers := `INSERT users (user_id, name, created_at, updated_at)
		  VALUES (@userID, @userName, @timestamp, @timestamp)`
		t := time.Now().Format("2006-01-02 15:04:05")
		params := map[string]interface{}{
			"userID":    u.userID,
			"userName":  u.userName,
			"timestamp": t,
		}
		stmtToUsers := spanner.Statement{
			SQL:    sqlToUsers,
			Params: params,
		}
		rowCountToUsers, err := txn.Update(ctx, stmtToUsers)
		_ = rowCountToUsers
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// add item specified item_id to specific user
func (d dbClient) addItemToUser(ctx context.Context, w io.Writer, u userParams, i itemParams) error {

	_, err := d.sc.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		sqlToUsers := `INSERT user_items (user_id, item_id, created_at, updated_at)
		  VALUES (@userID, @itemID, @timestamp, @timestamp)`
		t := time.Now().Format("2006-01-02 15:04:05")
		params := map[string]interface{}{
			"userID":    u.userID,
			"itemId":    i.itemID,
			"timestamp": t,
		}
		stmtToUsers := spanner.Statement{
			SQL:    sqlToUsers,
			Params: params,
		}
		rowCountToUsers, err := txn.Update(ctx, stmtToUsers)
		_ = rowCountToUsers
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// get what items the user has
func (d dbClient) userItems(ctx context.Context, w io.Writer, userID string) (*spanner.ReadOnlyTransaction, *spanner.RowIterator, error) {

	txn := d.sc.ReadOnlyTransaction()
	//defer txn.Close()
	sql := `select users.name,items.item_name,user_items.item_id from user_items 
		join items on items.item_id = user_items.item_id 
		join users on users.user_id = user_items.user_id 
		where user_items.user_id = @user_id`
	stmt := spanner.Statement{
		SQL: sql,
		Params: map[string]interface{}{
			"user_id": userID,
		},
	}

	iter := txn.Query(ctx, stmt)
	return txn, iter, nil

}

func (d dbClient) listItems(ctx context.Context) ([]itemParams, error) {

	txn := d.sc.ReadOnlyTransaction()
	defer txn.Close()
	sql := `select item_id,item_name from items`
	stmt := spanner.Statement{
		SQL: sql,
	}
	iter := d.sc.Single().Query(ctx, stmt)
	defer iter.Stop()

	items := []itemParams{}
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return items, nil
		}
		if err != nil {
			return items, err
		}
		var itemID, itemName string
		if err := row.Columns(&itemID, &itemName); err != nil {
			return items, err
		}
		items = append(items, itemParams{itemID: itemID})
	}
}
