package types

import (
	"sync"
)

// Create in-memory database storage for the websites
var webCollection map[string]string

// Structure to decode the JSON object [Not for instantiation]
type WebsiteRequest struct {
	URL string `json:"url"`
}

// Structure to store the response
type WebsiteResponse struct {
	URL       string `json:"url"`
	Available string `json:"inlist"`
	Active    string `json:"status"`
}

// create a new Website instance
func NewWebsiteResponse(URL string, Available string, Active string) *WebsiteResponse {
	return &WebsiteResponse{
		URL:       URL,
		Available: Available,
		Active:    Active,
	}
}

// Get the list of all the websites
func GetWebsitesList() []string {

	//Create a list to return the websites
	var websiteList []string

	// Iterate over the in-memory database
	for key := range webCollection {
		websiteList = append(websiteList, key)
	}

	return websiteList
}

// Add to the website in-memory storage
func UpdateWebsiteCollection(newWebsiteCollection []WebsiteRequest) {

	newTempCollection := map[string]string{}

	// Initialize the default state as "NIL"
	for _, item := range newWebsiteCollection {
		newTempCollection[item.URL] = "NIL"
	}

	webCollection = newTempCollection
}

// Update Website map by each goroutine
func UpdateWebsiteStatus(m *sync.Mutex, URL string, Active string) {
	//Updating the corresponding map value
	m.Lock()
	webCollection[URL] = Active
	m.Unlock()
}

// Get all the data from the map
func GetAllCollections() []WebsiteResponse {

	// Create a temporary instance
	var websiteCollectionResponse []WebsiteResponse

	for key := range webCollection {
		newWR := NewWebsiteResponse(key, "YES", webCollection[key])
		websiteCollectionResponse = append(websiteCollectionResponse, *newWR)
	}

	return websiteCollectionResponse
}

// Get the queried response from the map
func GetQueriedResponse(websiteURL string) *WebsiteResponse {

	// Perform a check for the presence of the website
	if _, ok := webCollection[websiteURL]; ok != true {
		return NewWebsiteResponse(websiteURL, "NO", "NIL")
	}
	return NewWebsiteResponse(websiteURL, "YES", webCollection[websiteURL])
}
