package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"spotdl/parser"
	"spotdl/provider"
)

func main() {
	searchTerms := "Lament Rain"
	maxResults := 15

	filePath, err := provider.YoutubeSearch(searchTerms, maxResults)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer os.Remove(*filePath)
	scripts, err := parser.ExtractScripts(*filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	script := parser.GetTargetJSVariable(regexp.MustCompile(`ytInitialData\s= {`), scripts)
	videos, err := parser.JavascriptDataProvider(script, "ytInitialData")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	j, err := json.Marshal(videos)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	os.WriteFile("videos.json", j, 0644)
}
