package searcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
	"ytvser/pkg/model"
)

type Searcher struct {
	SortBy string `json:"sort_by"` // 排序方式，可選 relevance、viewCount、date、rating
}

var ValidSortOptions = map[string]string{
	"relevance":   "EgIQAQ%3D%3D",
	"upload_date": "CAISAhAB",
	"views":       "CAMSAhAB",
	"rating":      "CAESAhAB",
}

func NewSearcher() Searcher {
	return Searcher{
		SortBy: "relevance", // 預設值為關聯性
	}
}

// IsValidSortOption 驗證給定的排序選項是否有效
func (s *Searcher) IsValidSortOption() bool {
	_, exists := ValidSortOptions[s.SortBy]
	return exists
}

// 設定排序選項
func (s *Searcher) SetSortOption(sortOption string) {
	s.SortBy = sortOption
	if !s.IsValidSortOption() {
		s.SortBy = "relevance"
		log.Fatal("Invalid sort option")
	}
}

func (s *Searcher) YoutubeSearch(searchTerms string, maxResults int) (results []model.Video, err error) {
	hyperText := new(string)
	baseURL := "https://www.youtube.com/results"
	params := url.Values{}
	params.Set("search_query", searchTerms)

	// Construct the search URL
	searchURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("http GET request failed %s: %w", searchURL, err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	*hyperText = fmt.Sprintf("%s/%s_%s.html", os.TempDir(), time.Now().Format("20060102150405"), searchTerms)
	err = os.WriteFile(*hyperText, body, 0644)

	defer os.Remove(*hyperText)
	scripts, err := ExtractScripts(*hyperText)

	if err != nil {
		return nil, fmt.Errorf("extractScripts Error: %w", err)
	}
	script := GetTargetJSVariable(regexp.MustCompile(`ytInitialData\s= {`), scripts)
	res, err := JavascriptDataProvider(script, "ytInitialData")

	if err != nil {
		return nil, fmt.Errorf("analize JS Error: %w", err)
	}

	return res, nil
}
