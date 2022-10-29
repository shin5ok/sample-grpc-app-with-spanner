package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"cloud.google.com/go/spanner"
	"github.com/stretchr/testify/assert"
)

var (
	fakeDbString = "projects/your-project-id/instances/foo-instance/databases/bar"
	fakeServing  Serving
)

/*
Note:
Before running test, run spanner emulator
*/
func init() {
	os.Setenv("SPANNER_EMULATOR_HOST", `localhost:9010`)
	ctx := context.Background()

	client, err := spanner.NewClient(ctx, fakeDbString)
	if err != nil {
		log.Fatal(err)
	}
	fakeServing = Serving{
		Client: dbClient{sc: client},
	}
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
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: %d. Got: %d", http.StatusOK, rr.Code)
	}

}
