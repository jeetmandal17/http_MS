package types

import(
	"sync"
)


// Create in-memory database storage for the websites
var webCollection map[string]bool

// Structure to decode the JSON object
type WebsiteRequest struct {
	URL string `json:"url"`
}

// Structure to store the respsonse
type WebsiteResponse struct {
	URL string		`json:"url"`
	Active bool		`json:"status"`
}

// create a new Website instance
func NewWebsiteResponse(URL string, Active bool) (*WebsiteResponse){
	return &WebsiteResponse{
		URL: URL,
		Active: Active,
	}
}

// Get the list of all the websites
func GetWebsitesList() ([]string){

	//Create a list to return the websites
	websiteList := []string{}

	for key := range webCollection{
		websiteList = append(websiteList, key)
	}

	return websiteList
}

// Add to the website inmemory storage
func UpdateWebsiteCollection(newWebsiteCollection []WebsiteRequest){

	newTempCollection := map[string]bool{}

	for _, item := range newWebsiteCollection{
		newTempCollection[item.URL] = false
	}

	webCollection = newTempCollection
}

// Update Website map by each goroutine 
func UpdateWebsiteStatus(m *sync.Mutex, URL string, Active bool){
	//Updating the corresponding map value
	m.Lock()
		webCollection[URL] = Active
	m.Unlock()
}

// Get all the data from the map
func GetAllCollections() ([]WebsiteResponse){

	// Create a temporary instance
	websiteCollectionResponse := []WebsiteResponse{}

	for key := range webCollection{
		newWR := NewWebsiteResponse(key, webCollection[key])
		websiteCollectionResponse = append(websiteCollectionResponse, *newWR)
	}

	return websiteCollectionResponse
}

// Get the queried response from the map
func GetQueriedResponse(websiteURL string) (*WebsiteResponse){
	return NewWebsiteResponse(websiteURL, webCollection[websiteURL])
}