package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/httpMS/Commons"
	"github.com/httpMS/handlers/types"
	"io"
	"log"
	"net/http"
	"os"
)

type UpdateList struct {
	logErrors *log.Logger
}

func NewUpdateList(log *log.Logger) *UpdateList {
	return &UpdateList{
		logErrors: log,
	}
}

func (u *UpdateList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Logging on the server side
	u.logErrors.Println("Update list for updating the list")

	// Logging on client side as ACK
	_, err := fmt.Fprint(rw, "Updating the list on server side")
	if err != nil {
		// Logging the error in the server logger
		u.logErrors.Println(Commons.ErrWritingOnClientSide, err)
	}

	// Handling the Update list
	updateData, err := io.ReadAll(r.Body)
	if err != nil {
		u.logErrors.Println(Commons.ErrReadingRequestBody, err)
	}

	// Create a WebsiteRequest array to store the values
	var webRequests []types.WebsiteRequest

	// Unmarshal the JSON object
	err = json.Unmarshal(updateData, &webRequests)
	if err != nil {
		u.logErrors.Println(Commons.ErrUnmarshalJSON, err)
	}

	// Update tge Website collection in the IN-MEMORY map
	types.UpdateWebsiteCollection(webRequests)

	// Start the monitoring service
	Commons.RoutineID = Commons.RoutineID + 1

	// Create a new logger for the new instance
	l := log.New(os.Stdout, "httpStatus-api", log.LstdFlags)

	// Get the instance from the interface
	httpStatusCheckerInstance := types.NewhttpChecker(l, Commons.RoutineID)
	go httpStatusCheckerInstance.InitializeMonitoring()
}
