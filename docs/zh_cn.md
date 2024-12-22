# Youtube 影片搜寻器

[English](https://github.com/miyago9267/ytdler/blob/master/README.md) | [繁體中文](https://github.com/miyago9267/ytdler/blob/master/docs/zh_tw.md) | 简体中文

本套件系 Youtube 影片搜寻器，为模拟 Youtube 影片搜寻之结果，可用于各种根据接受使用者搜寻之关键字返还结果之应用。

## 项目特色

- 无语言限制
- 支援多种排序方式：
  - 关联性 (`relevance`, 预设值)
  - 上传日期 (`upload_date`)
  - 观看次数 (`views`)
  - 评分 (`rating`)
- 提供简单易用的 API。

## 使用方式

### 安装

`go get github.com/miyago9267/ytdler`

### 使用

```go
import "github.com/miyago9267/ytdler"
```

### Model Schema

```go
type Video struct {
    ID          string \`json:"id"\`
    Title       string \`json:"title"\`
    Description string \`json:"description"\`
    Channel     string \`json:"channel"\`
    Duration    string \`json:"duration"\`
    Views       string \`json:"views"\`
    PublishTime string \`json:"publish_time"\`
    URL         string \`json:"url"\`
}
```

### 快速上手

[范例详见](https://github.com/miyago9267/ytdler/blob/master/cmd/main.go)

## License

本专案使用MIT协议进行开源，修改、分支、重制请遵守相同协议并标记原作者署名。

详细内容请参阅 [LICENSE](https://github.com/miyago9267/ytdler/blob/master/LICENSE)

## 贡献

本专案为 @aluo 开发并由 @miyago9267 维护为套件
