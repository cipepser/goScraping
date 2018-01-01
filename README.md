# goScraping

goでスクレイピングする時のライブラリである[goquery](https://github.com/PuerkitoBio/goquery)を触ってみたメモ

## 例題(main.go)

githubのレポジトリからstar数を抜き出すスクレイピング


## 触った途中のメモ

### Find("")の使い方

class要素に対しては"."オペレータでアクセスする。  
ブランクの手前まででOK


```go

s.Find("").Each(func(_ int, s *goquery.Selection) {
	fmt.Println(i, ": ", s.Text())
})

```


### Attributeを抜き出す

```go
s.Find("a").Each(func(_ int, s *goquery.Selection) {
	url, _ := s.Attr("href")
	fmt.Println(url)
})

```

### debugのとき

DOMのまま出力したほうが見やすい

``` go
html, _ := s.Html()
fmt.Println(html)
fmt.Println("--------------------------------")

data-typeは以下の形式でできるっぽい
s.Find("div[data-type='hogehoge']")

```

## References
* [goquery - a little like that j-thing, only in Go](https://github.com/PuerkitoBio/goquery)
* [goでスクレイピングするのにgoquery + bluemonday が最強な件](https://qiita.com/ryurock/items/b0466ad144f5e0555a95)

