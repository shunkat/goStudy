package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://ja.wikipedia.org/w/index.php?title=Category:%E6%B5%84%E5%9C%9F%E5%AE%97%E3%81%AE%E5%AF%BA%E9%99%A2&pagefrom=%E3%81%9F%E3%81%84%E3%81%AD%E3%82%93%E3%81%97%0A%E5%A4%A7%E5%BF%B5%E5%AF%BA+%28%E5%B9%B3%E5%A1%9A%E5%B8%82%29#mw-pages")
	if err != nil {
		panic("htmlの取得に失敗しました")
	}
	title := doc.Find(".mw-category a")
	title.Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())

	})
}
