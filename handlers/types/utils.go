package types

import (
	"fmt"
	"net/http"
	"sync"
)

func CheckWebsites() (){

	// Create a mutex for atomic operations
	mtx := new(sync.Mutex)

	// Get the list of all the websites
	var currentList = GetWebsitesList()

	// Instantiate all the individual goroutines
	for _, item := range currentList{
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
		UpdateWebsiteStatus(mtx, URL, false)
	}else{
		fmt.Println("succesfully pinged to the website")
		UpdateWebsiteStatus(mtx, URL, true)
	}
}