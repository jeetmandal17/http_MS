package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/httpMS/handlers/types"
)

type AllData struct{
	l *log.Logger
}

func NewAllData(log *log.Logger) (*AllData){
	return &AllData{
		l: log,
	}
}

// This is a GET request on the default path
func (a *AllData) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	
	// For returning all the data to the client
	a.l.Println("Handling the Access all data query")

	// Return data to the client
	websiteResponseList := types.GetAllCollections()

	// Convert in JSON object
	JSONOueryOutput, err := json.Marshal(websiteResponseList)

	if err != nil{
		fmt.Println("cannot marshal it into JSON")
	}

	// Convert into JSNO object
	fmt.Fprintf(rw, "Queried Website: %s", JSONOueryOutput)
}