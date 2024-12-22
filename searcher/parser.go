package searcher

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"youtube-searcher/model"

	"github.com/dop251/goja"
	"golang.org/x/net/html"
)

func javascriptDataProvider(script *string, variableName string) (videos []model.Video, err error) {
	vm := goja.New()
	if !strings.HasSuffix(*script, ";") {
		*script += ";"
	}
	*script += variableName + ";"
	// 執行 JavaScript 代碼
	value, err := vm.RunString(*script)
	if err != nil {
		fmt.Println("Error executing JS:", err)
		return
	}

	// 將結果轉換為 Go 的 interface{}
	rawData := value.Export()

	// 將數據轉換為 JSON 字串
	jsonData, err := json.MarshalIndent(rawData, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	data := model.YoutubeDataResult{}
	json.Unmarshal(jsonData, &data)

	for _, sectionContent := range data.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, videoContent := range sectionContent.ItemSectionRenderer.Contents {
			videoID := videoContent.VideoRenderer.VideoID
			if videoID == "" {
				continue
			}
			video := model.Video{
				ID:          videoID,
				Views:       videoContent.VideoRenderer.ViewCountText.SimpleText,
				URL:         fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID),
				Duration:    videoContent.VideoRenderer.LengthText.SimpleText,
				PublishTime: videoContent.VideoRenderer.PublishedTimeText.SimpleText,
			}
			if len(videoContent.VideoRenderer.Title.Runs) > 0 {
				video.Title = videoContent.VideoRenderer.Title.Runs[0].Text
			}
			if len(videoContent.VideoRenderer.DetailedMetadataSnippets) > 0 {
				video.Description = ""
				for _, description := range videoContent.VideoRenderer.DetailedMetadataSnippets[0].SnippetText.Runs {
					video.Description += description.Text + " "
				}
			}
			if len(videoContent.VideoRenderer.LongBylineText.Runs) > 0 {
				video.Channel = videoContent.VideoRenderer.LongBylineText.Runs[0].Text
			}
			videos = append(videos, video)
		}
	}
	return
}

func getTargetJSVariable(target *regexp.Regexp, scripts []string) (script *string) {
	script = new(string)
	// re := regexp.MustCompile(`ytInitialData\s= {`)
	index := 0
	for i, script := range scripts {
		// fmt.Printf("Script %d:\n%s\n\n", i+1, script)
		if target.MatchString(script) {
			log.Println("Found target JS variable")
			index = i
			break
		}
	}
	*script = scripts[index]
	return
}

func extractScripts(htmlBody string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var scripts []string

	// 遞迴遍歷 HTML 節點
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			if n.FirstChild != nil {
				scripts = append(scripts, strings.TrimSpace(n.FirstChild.Data))
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return scripts, nil
}
