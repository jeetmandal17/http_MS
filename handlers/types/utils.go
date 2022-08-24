package types

import (
	"fmt"
	"github.com/httpMS/Commons"
	"log"
	"net/http"
	"sync"
	"time"
)

// Implement an interface to encapsulate the website Checker
type StatusChecker interface {
	InitializeMonitoring(RoutineId int)
}

// An instance for Monitoring service
type httpChecker struct {
	logEvents *log.Logger
	RoutineID int
}

// Create an instance of the Service
func NewhttpChecker(logEvents *log.Logger, RoutineID int) *httpChecker {
	return &httpChecker{
		logEvents: logEvents,
		RoutineID: RoutineID,
	}
}

// This is the required function for monitoring
func (h httpChecker) InitializeMonitoring() {

	// Ping the websites periodically
	// Launch the CheckWebsites service

	// checking how many go routines are running
	//fmt.Println(h.RoutineID)
	//fmt.Println(runtime.NumGoroutine())
	for {
		// check if there is a new Routine or not
		if h.RoutineID != Commons.RoutineID {
			// Stop this instance
			fmt.Println(h.RoutineID, " - ", Commons.RoutineID)
			break
		}

		// Check for the listed websites
		h.logEvents.Println("")
		CheckWebsites()

		// Check for the websites every 10 seconds
		time.Sleep(60 * time.Second)
	}

	// Close this task and record the event
	h.logEvents.Println("Dumping the Old Status Checker Instance with ID: ", h.RoutineID)
}

// This function is used to ping the required server
func CheckWebsites() {

	// Create a mutex for atomic operations
	mtx := new(sync.Mutex)

	// Get the list of all the websites
	var currentList = GetWebsitesList()

	// Instantiate all the individual goroutines
	for _, item := range currentList {
		go PingURL(mtx, item)
	}
}

// Function to deploy multiple goroutine to ping to the server
func PingURL(mtx *sync.Mutex, URL string) {

	// Create pinging instances from the
	requestURL := "https://" + URL
	_, err := http.Get(requestURL)

	if err != nil {
		fmt.Println("failed to ping to the website")
		// Edit the current website status IN-Memory
		UpdateWebsiteStatus(mtx, URL, "DOWN")
	} else {
		fmt.Println("successfully pinged to the website")
		UpdateWebsiteStatus(mtx, URL, "UP")
	}
}
