package main

import (
	"context"
	"encoding/json"
	"strings"

	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly/v2"
)

type Input struct {
	Id string `json:"id"`
}

type Output struct {
	Id     string `json:"id"`
	Exists bool   `json:"exists"`
}

func HandleRequest(ctx context.Context, event json.RawMessage) (string, error) {
	var input Input
	if err := json.Unmarshal(event, &input); err != nil {
		return "", err
	}
	var output Output

	// コレクターを作成
	c := colly.NewCollector(
		colly.AllowedDomains("www.db.yugioh-card.com"),
		colly.DetectCharset(),
	)

	// カードが存在しない場合
	c.OnHTML("div[class=no_data]", func(e *colly.HTMLElement) {
		output.Id = input.Id
		output.Exists = false

		fmt.Println("カードが存在していません:", output.Id)
	})

	// カードが存在する場合
	c.OnHTML("div[id=cardname]", func(e *colly.HTMLElement) {
		rawContents := strings.Split(strings.Replace(e.DOM.Text(), "\t", "", -1), "\n")

		var contents []string
		for _, content := range rawContents {
			if len(strings.TrimSpace(content)) > 0 {
				contents = append(contents, strings.TrimSpace(content))
			}
		}

		fmt.Println("カードが存在しています:", contents[1])
		output.Id = input.Id
		output.Exists = true
	})

	// ニューロンにアクセス
	c.Visit("https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=" + input.Id + "&request_locale=ja")

	// JSONを返却
	json, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func main() {
	lambda.Start(HandleRequest)
}
