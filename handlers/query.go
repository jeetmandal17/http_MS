package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/httpMS/handlers/types"
)

type Query struct{
	l *log.Logger
}

func NewQuery(log *log.Logger) (*Query){
	return &Query{
		l: log,
	}
}

func (q *Query) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	q.l.Println("Get the Queried website")
	
	// Handling the GET request for queried data
	QueryData := r.FormValue("website")

	// Get the Queried data from the in-memory getter
	websiteResponse := types.GetQueriedResponse(QueryData)

	// Convert in JSON object
	JSONOueryOutput, err := json.Marshal(websiteResponse)

	if err != nil{
		fmt.Println("cannot marshal it into JSON")
	}

	// Convert into JSNO object
	fmt.Fprintf(rw, "Queried Website: %s", JSONOueryOutput)
}