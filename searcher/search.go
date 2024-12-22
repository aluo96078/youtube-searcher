package searcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"youtube-searcher/model"
)

// Searcher 定義搜尋器的結構
// SortBy 支援以下排序方式：
// - relevance: 關聯性 (預設值)
// - upload_date: 上傳日期
// - views: 觀看次數
// - rating: 評分
type Searcher struct {
	SortBy string `json:"sort_by"`
}

// ValidSortOptions 定義有效的排序選項
var ValidSortOptions = map[string]string{
	"relevance":   "EgIQAQ%3D%3D",
	"upload_date": "CAISAhAB",
	"views":       "CAMSAhAB",
	"rating":      "CAESAhAB",
}

// NewSearcher 創建並初始化一個 Searcher
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

// GetSortParameter 取得對應的排序參數 (sp 值)
func (s *Searcher) GetSortParameter() string {
	if param, exists := ValidSortOptions[s.SortBy]; exists {
		return param
	}
	return ValidSortOptions["relevance"] // 預設為關聯性
}

// SetSortBy 設定排序方式，若無效則回退為預設值
func (s *Searcher) SetSortBy(sortBy string) {
	if _, exists := ValidSortOptions[sortBy]; exists {
		s.SortBy = sortBy
	} else {
		s.SortBy = "relevance"
	}
}

func (s *Searcher) Search(keyword string, maxResults int) (videos []model.Video, err error) {
	baseURL := "https://www.youtube.com/results"
	params := url.Values{}
	params.Set("search_query", keyword)
	params.Set("sp", s.GetSortParameter())

	// Construct the search URL
	searchURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(searchURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	reg, err := regexp.Compile(`ytInitialData\s= {`)
	if err != nil {
		return
	}
	scripts, err := extractScripts(string(body))
	script := getTargetJSVariable(reg, scripts)
	return javascriptDataProvider(script, "ytInitialData")
}
