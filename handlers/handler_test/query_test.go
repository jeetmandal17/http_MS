package handler_test

import (
	"github.com/httpMS/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestQueryAPI(t *testing.T) {

	// Create a request to the API endpoint
	request, err := http.NewRequest("GET", "/GET?website=www.xyzz.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Creates a recorder to receive the request from the endpoint
	rr := httptest.NewRecorder()
	l := log.New(os.Stdout, "QueryAPI-Test", log.LstdFlags)

	// Create a new handler
	getQueryDataHandler := handlers.NewQuery(l)

	// Send the package to the API endpoint
	getQueryDataHandler.ServeHTTP(rr, request)

	// Now check the rr code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error connecting to the client")
	}

	// Check the body received in the response
	if rr.Body.String() == "" {
		t.Errorf("Received unexpected response from the server %v", rr.Body.String())
	}
}
