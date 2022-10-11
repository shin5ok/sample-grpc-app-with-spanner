package main

import (
	"context"
	"io"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

type GameUserOperation interface {
	createUser(io.Writer, *spanner.Client, userParams) error
	updateScore(io.Writer, *spanner.Client, userParams, int64) error
}

type userParams struct {
	userID   string
	userName string
}

type dbClient struct{}

func spannerNewClient(dbString string) (*spanner.Client, error) {

	ctx := context.Background()
	client, err := spanner.NewClient(ctx, dbString)
	if err != nil {
		return &spanner.Client{}, err
	}
	return client, nil
}

func (d dbClient) createUser(w io.Writer, client *spanner.Client, u userParams) error {

	ctx := context.Background()
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
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

		sqlToScores := `INSERT scores (user_id, score_id, score, created_at, updated_at)
		  VALUES (@userID, @scoreID, 0, @timestamp, @timestamp)`
		stmtToScores := spanner.Statement{SQL: sqlToScores, Params: map[string]interface{}{}}
		scoreID, _ := uuid.NewUUID()
		stmtToScores.Params["scoreID"] = scoreID.String()
		stmtToScores.Params["userID"] = u.userID
		stmtToScores.Params["timestamp"] = t

		rowCountToScores, err := txn.Update(ctx, stmtToScores)
		_ = rowCountToScores
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (d dbClient) updateScore(w io.Writer, client *spanner.Client, u userParams, score int64) error {
	ctx := context.Background()
	_, err := client.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		sqlToScore := `update scores set score = @newScore, updated_at = @timestamp where user_id = (select user_id from users where name = @name limit 1)`
		t := time.Now().Format("2006-01-02 15:04:05")
		params := map[string]interface{}{
			"name":      u.userName,
			"timestamp": t,
			"newScore":  score,
		}
		stmtToScore := spanner.Statement{
			SQL:    sqlToScore,
			Params: params,
		}

		_, err := txn.Update(ctx, stmtToScore)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
