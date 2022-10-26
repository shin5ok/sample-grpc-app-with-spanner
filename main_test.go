package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spannertest"
	"cloud.google.com/go/spanner/spansql"
	"github.com/stretchr/testify/assert"
)

var (
	client       *spanner.Client
	fakeDbString = "projects/your-project-id/instances/foo-instance/databases/bar"
	fakeServing  = Serving{}
)

func init() {
	srv, err := spannertest.NewServer("localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("SPANNER_EMULATOR_HOST", srv.Addr)
	// os.Setenv("PORT", "12820")
	ctx := context.Background()

	client, err = spanner.NewClient(ctx, fakeDbString)
	if err != nil {
		log.Fatal(err)
	}
	fakeServing = Serving{
		Client: dbClient{sc: client},
	}

	var list []spansql.DDLStmt

	schemas, _ := filepath.Glob("schemas/*.sql")
	for _, file := range schemas {
		buf, _ := os.ReadFile(file)
		stmt, err := spansql.ParseDDLStmt(string(buf))
		if err != nil {
			log.Print("parse error", err)
		}
		list = append(list, stmt)
	}

	sqlDDL := spansql.DDL{
		List: list,
	}
	_ = sqlDDL
	// err = srv.UpdateDDL(&sqlDDL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func Test_run(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fakeServing.pingPong)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: %d. Got: %d", http.StatusOK, rr.Code)
	}

}

func Test_createUser(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/users/test-user", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fakeServing.createUser)
	assert.NotNil(t, handler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: %d. Got: %d", http.StatusOK, rr.Code)
	}

}
