package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/httpMS/Commons"
	"log"
	"net/http"

	"github.com/httpMS/handlers/types"
)

type Query struct {
	l *log.Logger
}

func NewQuery(log *log.Logger) *Query {
	return &Query{
		l: log,
	}
}

func (q *Query) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Logging on the server side
	q.l.Println("Get the Queried website")

	// Logging on client side as aCK
	_, err := fmt.Fprint(rw, "Sending the Queried website Status")
	if err != nil {
		q.l.Println(Commons.ErrWritingOnClientSide, err)
	}

	// Handling the GET request for queried data
	QueryData := r.FormValue("website")

	// Get the Queried data from the in-memory getter
	websiteResponse := types.GetQueriedResponse(QueryData)

	// Convert in JSON object
	JSONQueryOutput, err := json.Marshal(websiteResponse)
	if err != nil {
		q.l.Println(Commons.ErrMarshalJSON, err)
	}

	// Write the JSON object on client side
	fmt.Fprintf(rw, "Queried Website: %s", JSONQueryOutput)
}
