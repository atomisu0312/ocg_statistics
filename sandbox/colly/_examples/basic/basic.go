package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.db.yugioh-card.com"),
		colly.DetectCharset(),
	)

	c.OnHTML("div[id=cardname]", func(e *colly.HTMLElement) {
		rawContents := strings.Split(strings.Replace(e.DOM.Text(), "\t", "", -1), "\n")

		var contents []string
		for _, content := range rawContents {
			if len(strings.TrimSpace(content)) > 0 {
				contents = append(contents, strings.TrimSpace(content))
			}
		}

		fmt.Println("フィルタリング後タイトル:", contents[2])

	})

	c.OnHTML("div[class=CardText] > div[class=item_box_text]", func(e *colly.HTMLElement) {
		rawContents := strings.Split(strings.Replace(e.DOM.Text(), "\t", "", -1), "\n")

		var contents []string
		for _, content := range rawContents {
			if len(strings.TrimSpace(content)) > 0 {
				contents = append(contents, strings.TrimSpace(content))
			}
		}

		fmt.Println("フィルタリング後:", contents[1])

	})

	c.Visit("https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=4817&request_locale=ja")
}
