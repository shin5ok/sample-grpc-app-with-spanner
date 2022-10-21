package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type GameUserOperation interface {
	createUser(context.Context, io.Writer, userParams) error
	addItemToUser(context.Context, io.Writer, userParams, itemParams) error
	listUsers(context.Context, io.Writer, string) ([]map[string]interface{}, error)
}

type userParams struct {
	userID   string
	userName string
}

type itemParams struct {
	itemID    string
	itemPrice int64
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

// create a user while initializing score field
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

// update score field corresponding to specified user
func (d dbClient) listUsers(ctx context.Context, w io.Writer, name string) ([]map[string]interface{}, error) {
	txn := d.sc.ReadOnlyTransaction()
	defer txn.Close()
	sql := "SELECT users.user_id,users.name from users join user_items on users.user_id = user_items.user_id where users.name like @name;"
	stmt := spanner.Statement{
		SQL: sql,
		Params: map[string]interface{}{
			"name": fmt.Sprintf("%s%%", name),
		},
	}

	iter := txn.Query(ctx, stmt)
	defer iter.Stop()

	results := []map[string]interface{}{}
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return results, err
		}
		var userName string
		var userId string
		//var userScore int64
		//if err := row.Columns(&userName, &userId, &userScore); err != nil {
		if err := row.Columns(&userName, &userId); err != nil {
			return results, err
		}

		results = append(results, map[string]interface{}{"name": userName, "id": userId})

	}

	return results, nil
}
