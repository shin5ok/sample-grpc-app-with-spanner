package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/spannertest"
	"github.com/stretchr/testify/assert"
)

func Test_run(t *testing.T) {

	ctx := context.Background()

	srv, err := spannertest.NewServer("localhost:0")
	assert.Nil(t, err)
	os.Setenv("SPANNER_EMULATOR_HOST", srv.Addr)
	os.Setenv("PORT", "12820")
	fakeDbString := "projects/your-project-id/instances/foo/databases/bar"
	client, err := spanner.NewClient(ctx, fakeDbString)
	assert.Nil(t, err)

	s := Serving{
		Client: dbClient{sc: client},
	}

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.pingPong)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: %d. Got: %d", http.StatusOK, rr.Code)
	}

}
