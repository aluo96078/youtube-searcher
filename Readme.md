# Youtube 影片列表搜尋

`youtube-searcher` 是一個用於搜尋影片的 Go 語言庫，專為 YouTube 搜尋設計，並支援多種排序方式。此庫提供了一種簡單的方法來執行搜尋，並根據關聯性、上傳日期、觀看次數或評分進行排序。

---

## 功能特色

- 支援多種排序方式：
  - 關聯性 (`relevance`, 預設值)
  - 上傳日期 (`upload_date`)
  - 觀看次數 (`views`)
  - 評分 (`rating`)
- 自動驗證排序選項是否有效。
- 提供簡單易用的 API。

---

## 安裝

確保您的專案已經初始化 Go modules，然後執行以下指令安裝此套件：

```bash
go get github.com/aluo96078/youtube-searcher
```

---

## 快速入門

以下是一個使用 `searcher` 套件的基本範例：

```go
package main

import (
	"fmt"
	"log"
	"searcher"
)

func main() {
	// 創建一個新的搜尋器實例
	s := searcher.NewSearcher()

	// 設定排序方式為觀看次數
	s.SetSortBy("views")

	// 確認排序方式是否有效
	if !s.IsValidSortOption() {
		log.Fatalf("無效的排序方式：%s", s.SortBy)
	}

	// 搜尋關鍵字
	videos, err := s.Search("Golang 教學", 10)
	if err != nil {
		log.Fatalf("搜尋失敗: %v", err)
	}

	// 打印搜尋結果
	for _, video := range videos {
		fmt.Printf("影片標題: %s, 影片網址: %s\n", video.Title, video.URL)
	}
}
```

---

## API 文件

### 類別: `Searcher`

#### 1. **屬性**

- `SortBy`: (字串) 指定搜尋排序方式。

#### 2. **方法**

##### (1) `NewSearcher() Searcher`
創建並初始化一個 `Searcher` 實例，預設排序方式為 `relevance`。

##### (2) `SetSortBy(sortBy string)`
設定排序方式。如果指定的排序方式無效，將回退為預設值 `relevance`。

##### (3) `IsValidSortOption() bool`
檢查目前的排序方式是否有效。

##### (4) `GetSortParameter() string`
返回對應排序方式的 YouTube API 參數值。

##### (5) `Search(keyword string, maxResults int) ([]model.Video, error)`
執行搜尋：
- `keyword`: 搜尋的關鍵字。
- `maxResults`: 最大搜尋結果數量。

返回一個包含 `model.Video` 結構的切片，或返回錯誤。

---

## 結構

### `model.Video`

此結構用於表示影片資訊：

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

---

## 支援的排序選項

排序方式的有效選項如下：

| 排序方式       | 描述   |
|----------------|--------|
| `relevance`    | 關聯性 |
| `upload_date`  | 上傳日期 |
| `views`        | 觀看次數 |
| `rating`       | 評分   |

---

## 錯誤處理

當發生以下情況時，`youtube-searcher` 的方法可能返回錯誤：

1. HTTP 請求失敗。
2. 無法解析搜尋結果頁面。
3. 給定的排序方式無效。

---

## 注意事項

1. 此庫使用正則表達式解析 YouTube 頁面資料，可能需要根據 YouTube 的更新進行維護。
2. 在執行高頻率請求時，請注意遵守 YouTube 的使用條款與限制。

---

## 貢獻

歡迎提交 Issue 或 Pull Request！請確保您的代碼通過了所有測試並符合項目代碼風格。

---

## 授權

此專案採用 [MIT 授權](LICENSE)。

