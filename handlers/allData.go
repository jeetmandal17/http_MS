package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/httpMS/Commons"
	"log"
	"net/http"

	"github.com/httpMS/handlers/types"
)

type AllData struct {
	logErrors *log.Logger
}

func NewAllData(log *log.Logger) *AllData {
	return &AllData{
		logErrors: log,
	}
}

// This is a GET request on the default path
func (a *AllData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// Logging on server side
	a.logErrors.Println("Handling the Access all data query")

	// Logging on client side to get an ACK
	_, err := fmt.Fprint(rw, "Sending all the data to client")
	if err != nil {
		a.logErrors.Println(Commons.ErrWritingOnClientSide, err)
	}
	// Return data to the client
	websiteResponseList := types.GetAllCollections()

	// Convert in JSON object
	JSONQueryOutput, err := json.Marshal(websiteResponseList)
	if err != nil {
		a.logErrors.Println(Commons.ErrMarshalJSON, err)
	}

	// Convert into JSON object
	fmt.Fprintf(rw, "Queried Website: %s", JSONQueryOutput)
}
