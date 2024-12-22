package model

// YoutubeDataResult 代表 YouTube 的搜尋結果資料結構
type YoutubeDataResult struct {
	Contents Contents `json:"contents"` // 頂層內容
}

// Contents 頂層內容
type Contents struct {
	TwoColumnSearchResultsRenderer TwoColumnSearchResultsRenderer `json:"twoColumnSearchResultsRenderer"`
}

// TwoColumnSearchResultsRenderer 雙欄式搜尋結果渲染器
type TwoColumnSearchResultsRenderer struct {
	PrimaryContents PrimaryContents `json:"primaryContents"`
}

// PrimaryContents 主要內容區塊
type PrimaryContents struct {
	SectionListRenderer SectionListRenderer `json:"sectionListRenderer"`
}

// SectionListRenderer 區段列表渲染器
type SectionListRenderer struct {
	Contents []SectionContent `json:"contents"` // 區段內容陣列
}

// SectionContent 區段內容
type SectionContent struct {
	ItemSectionRenderer ItemSectionRenderer `json:"itemSectionRenderer"`
}

// ItemSectionRenderer 單項區塊渲染器
type ItemSectionRenderer struct {
	Contents []VideoContent `json:"contents"` // 每項內容
}

// VideoContent 影片內容
type VideoContent struct {
	VideoRenderer VideoRenderer `json:"videoRenderer"`
}

// VideoRenderer 影片渲染器
type VideoRenderer struct {
	Avatar                   Avatar                     `json:"avatar"`
	DetailedMetadataSnippets []DetailedMetadataSnippets `json:"detailedMetadataSnippets"`
	LengthText               LengthText                 `json:"lengthText"`
	Title                    Title                      `json:"title"`
	ViewCountText            ViewCountText              `json:"viewCountText"`
	VideoID                  string                     `json:"videoId"`
	Thumbnail                Thumbnail                  `json:"thumbnail"`
	LongBylineText           LongBylineText             `json:"longBylineText"`
	PublishedTimeText        PublishedTimeText          `json:"publishedTimeText"`
}

type PublishedTimeText struct {
	SimpleText string `json:"simpleText"`
}

// DetailedMetadataSnippets 詳細元數據片段
type DetailedMetadataSnippets struct {
	MaxOneLine  bool `json:"maxOneLine"`
	SnippetText struct {
		Runs []TextRun `json:"runs"`
	} `json:"snippetText"`
}

// Avatar 頻道頭像
type Avatar struct {
	DecoratedAvatarViewModel DecoratedAvatarViewModel `json:"decoratedAvatarViewModel"`
}

// DecoratedAvatarViewModel 裝飾過的頭像視圖模型
type DecoratedAvatarViewModel struct {
	A11YLabel string          `json:"a11yLabel"` // 輔助說明標籤
	Avatar    AvatarViewModel `json:"avatar"`
}

// AvatarViewModel 頭像視圖模型
type AvatarViewModel struct {
	AvatarImageSize string `json:"avatarImageSize"` // 頭像圖片大小
	Image           Image  `json:"image"`           // 圖片資訊
}

// Image 圖片結構
type Image struct {
	Sources []ImageSource `json:"sources"` // 圖片來源列表
}

// ImageSource 單個圖片來源
type ImageSource struct {
	Height int    `json:"height"` // 圖片高度
	URL    string `json:"url"`    // 圖片 URL
	Width  int    `json:"width"`  // 圖片寬度
}

// LengthText 影片長度資訊
type LengthText struct {
	Accessibility Accessibility `json:"accessibility"`
	SimpleText    string        `json:"simpleText"` // 影片長度文字
}

// Accessibility 輔助說明結構
type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

// AccessibilityData 輔助說明的數據
type AccessibilityData struct {
	Label string `json:"label"` // 輔助說明文字
}

// Title 影片標題
type Title struct {
	Accessibility Accessibility `json:"accessibility"`
	Runs          []TextRun     `json:"runs"`
}

// TextRun 單個文字段落
type TextRun struct {
	Text string `json:"text"` // 文字內容
}

// ViewCountText 總觀看次數資訊
type ViewCountText struct {
	SimpleText string `json:"simpleText"` // 觀看次數的文字表示
}

// Thumbnail 縮略圖資訊
type Thumbnail struct {
	Thumbnails []ImageSource `json:"thumbnails"` // 縮略圖列表
}

// LongBylineText 頻道名稱與相關資訊
type LongBylineText struct {
	Runs []BylineRun `json:"runs"`
}

// BylineRun 單個頻道名稱的段落
type BylineRun struct {
	Text               string             `json:"text"` // 頻道名稱
	NavigationEndpoint NavigationEndpoint `json:"navigationEndpoint"`
}

// NavigationEndpoint 導航端點
type NavigationEndpoint struct {
	BrowseEndpoint BrowseEndpoint `json:"browseEndpoint"`
}

// BrowseEndpoint 瀏覽端點
type BrowseEndpoint struct {
	BrowseID string `json:"browseId"` // 頻道的瀏覽 ID
}
