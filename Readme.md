# YouTube Video Searcher

English | [繁體中文](https://github.com/miyago9267/ytdler/blob/master/docs/zh_tw.md) | [简体中文](https://github.com/miyago9267/ytdler/blob/master/docs/zh_cn.md)

This package as a YouTube video searcher, simulating YouTube video search results. It can be used for applications that return results based on user-provided keywords.

## Features

- No language restrictions
- Supports multiple sorting options:
  - Relevance (default)
  - Upload date
  - Views
  - Rating
- Simple and easy-to-use API

## Getting Started

### Installation

```bash
go get github.com/miyago9267/ytdler
```

### Usage

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

### Usage Sample

[See Example](https://github.com/miyago9267/ytdler/blob/main/cmd/main.go)

## License

This project is open-sourced under the MIT license. For modifications, branching, and reproduction, please comply with the same license and credit the original author.

For more details, refer to LICENSE.

## Contributions

This project was developed by @aluo and maintained as a package by @miyago9267.
