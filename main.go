package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	URL := "https://github.com/cipepser/goNLP100knock2015"

	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}
	p := bluemonday.NewPolicy()
	p.AllowElements("a")

	doc.Find("div > div > div > div > div > ul > li > a.social-count").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			fmt.Println("star:", strings.TrimSpace(s.Text()))
		}
	})

}
