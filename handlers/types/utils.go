package types

import (
	"fmt"
	"net/http"
)

func CheckWebsites() (){

	// Get the list of all the websites 
	currentList := GetWebsitesList()

	for _, item := range currentList{

		
	}
}

func PingURL(URL string) (){

	// Ping the URL every 1 minute
	// Create pinging instances from the 
	for {
		// Create pinging instances from the 
		requestURL := "https://" + URL
		httpClient, err := http.Get(requestURL)

		if err != nil {
			fmt.Println("")
		}
	}

	

}