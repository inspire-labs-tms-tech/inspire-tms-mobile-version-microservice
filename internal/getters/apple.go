package getters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type iTunesResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		Version string `json:"version"`
	} `json:"results"`
}

const url = "https://itunes.apple.com/lookup?bundleId=com.inspiretmstech.mobile"

func main() {

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the JSON response
	var data iTunesResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Check if results exist and print the version
	if data.ResultCount > 0 && len(data.Results) > 0 {
		fmt.Println("Current Version:", data.Results[0].Version)
	} else {
		fmt.Println("No results found for the given bundle ID.")
	}
}
