package provider

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func YoutubeSearch(searchTerms string, maxResults int) (outputFilePath *string, err error) {
	outputFilePath = new(string)
	baseURL := "https://www.youtube.com/results"
	params := url.Values{}
	params.Set("search_query", searchTerms)

	// Construct the search URL
	searchURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(searchURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	*outputFilePath = fmt.Sprintf("%s/%s_%s.html", os.TempDir(), time.Now().Format("20060102150405"), searchTerms)
	err = os.WriteFile(*outputFilePath, body, 0644)
	return
}
