package main

import (
	"fmt"

	"github.com/miyago9267/ytvser/pkg/searcher"
)

func main() {
	searchTerms := "Golang土撥鼠"
	maxResults := 15

	// 創建 Searcher 物件
	s := searcher.NewSearcher()

	// s.SetSortOption("views") // 可設定排序順序

	// 執行搜尋並取得結果
	result, err := s.YoutubeSearch(searchTerms, maxResults)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 印出搜尋結果
	for _, item := range result {
		fmt.Println(
			"Video Title:", item.Title,
			"Video URL:", item.URL,
			"Duration:", item.Duration,
			"Views:", item.Views)
	}
}
