package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/httpMS/handlers/types"
)

type UpdateList struct{
	l *log.Logger
}

func NewUpdateList(log *log.Logger) (*UpdateList){
	return &UpdateList{
		l: log,
	}
}

func (u *UpdateList) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	u.l.Println("Update list for updating the list")

	// Handling the Update list
	updateData, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in getting the update values")
	}

	// Create a WebsiteRequest array to store the values
	webRequests := []types.WebsiteRequest{}

	// Unmarshal the JSON object
	err = json.Unmarshal(updateData, &webRequests)
	if err != nil {
		fmt.Println("cannot unmarshal the data into struct", err)
	}

	// Update tge Website collection in the IN-MEMORY map
	types.UpdateWebsiteCollection(webRequests)

	// Launch the CheckWebites service
	for {
		types.CheckWebsites()

		// Check for the websites every 10 seconds
		time.Sleep(10*time.Second)
	}
}