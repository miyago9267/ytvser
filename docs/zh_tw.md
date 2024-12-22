# Youtube 影片搜尋器

[English](https://github.com/miyago9267/ytdler/blob/master/README.md) | 繁體中文 | [简体中文](https://github.com/miyago9267/ytdler/blob/master/docs/zh-cn.md)

本套件係 Youtube 影片搜尋器，為模擬 Youtube 影片搜尋之結果，可用於各種根據接受使用者搜尋之關鍵字返還結果之應用。

## 專案特色

- 無語言限制
- 支援多種排序方式：
  - 關聯性 (`relevance`, 預設值)
  - 上傳日期 (`upload_date`)
  - 觀看次數 (`views`)
  - 評分 (`rating`)
- 提供簡單易用的 API。

## 使用方式

### 安裝

`go get github.com/miyago9267/ytdler`

### 使用

```go
import "github.com/miyago9267/ytdler"
```

### Model Schema

```go
type Video struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Channel     string `json:"channel"`
    Duration    string `json:"duration"`
    Views       string `json:"views"`
    PublishTime string `json:"publish_time"`
    URL         string `json:"url"`
}
```

### 快速上手

[範例詳見](https://github.com/miyago9267/ytdler/blob/master/cmd/main.go)

## License

本專案使用MIT協議進行開源，修改、分支、重製請遵守相同協議並標記原作者署名。

詳細內容請參閱 [LICENSE](https://github.com/miyago9267/ytdler/blob/master/LICENSE)

## 貢獻

本專案為 @aluo 開發並由 @miyago9267 維護為套件
