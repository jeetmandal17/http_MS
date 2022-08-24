package handler_test

import (
	"github.com/httpMS/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUpdateList(t *testing.T) {

	// Generate a request to be
	request, err := http.NewRequest("POST", "/POST", nil)
	if err != nil {
		t.Fatal("cannot connect to endpoint")
	}

	// Creating http recorder
	rr := httptest.NewRecorder()
	l := log.New(os.Stdout, "UpdateList-Test", log.LstdFlags)

	// Create a handler instance
	getUpdateListHandler := handlers.NewQuery(l)

	// Send the package to the API endpoint
	getUpdateListHandler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Cannot connect to the endpoint")
	}
}
