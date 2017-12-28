package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	URL := "https://github.com/cipepser/goNLP100knock2015"

	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li").Each(func(i int, s *goquery.Selection) {

		// elem := s.Find("thead > td").Text()

		p := bluemonday.UGCPolicy()
		p.AllowElements("form")
		html := p.Sanitize(s.Text())

		fmt.Println(html)

		// elem := s.Find("button").Text()
		// fmt.Println(i, ": ", elem)
	})

}
