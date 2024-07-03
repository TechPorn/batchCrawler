package parser

import (
	"batchCrawler/engine"
	"regexp"
	"time"
)

var provinceRe = regexp.MustCompile(`(?s)<a href="/content.php\?cid=([0-9]+)" target="_blank">.*?<div class="Title">([^<]+)</div>.*?<span>([^<]+)</span>`)

func ParseProvince(contents string) engine.ParseResult {
	currentDate := time.Now()
	today := currentDate.Format("2006-01-02")

	provinceMatches := provinceRe.FindAllStringSubmatch(contents, -1)

	result := engine.ParseResult{}
	if len(provinceMatches) > 0 {
		for _, m := range provinceMatches {
			title := m[2]
			id := m[1]
			date := m[3]
			url := "https://www.lfgvip.com/content.php?cid=" + id
			if date == today {
				result.Requests = append(result.Requests, engine.Request{
					Url: url,
					ParserFunc: func(s string) engine.ParseResult {
						return ParseProfile(s, title, url, id, date)
					},
				})
			}
		}

	}

	return result
}
