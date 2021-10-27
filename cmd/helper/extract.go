package helper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
)

type Content struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func ExtractContent(selector string, body io.Reader) []Content {
	var content []Content
	doc, err := goquery.NewDocumentFromReader(body)
	PrintIfError(err)
	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		link, exist := selection.Find("a").Attr("href")
		if exist == false {
			fmt.Println("attribute not exists")
			panic(errors.New("attribute is not exists"))
		}
		data := Content{
			Title: selection.Find("a").Text(),
			Link:  link,
		}
		content = append(content, data)
	})
	return content
}

