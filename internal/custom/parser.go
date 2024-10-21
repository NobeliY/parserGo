package custom

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func parseMovies(g *geziyor.Geziyor, r * client.Response) {
	r.HTMLDoc.Find("div.shedule").Each(func(i int, s *goquery.Selection) {
		var sessions = strings.Split(s.Find(".schedule").Text(), "\n")
		sessions = sessions[:len(sessions)-1]
		for i := 0; i < len(sessions); i++ {
			sessions[i] = strings.Trim(sessions[i], "\n ")
		}
	})
}

func ParseUrls(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.catalog_top_sections").Each(Catalog)
}

func Catalog(i int, s *goquery.Selection) {
	var items = s.Find(".catalog_top_sections__item")
	fmt.Println(items)
	var sessions = strings.Split(s.Find(".catalog_top_sections__item").Text(), "\n")
	sessions = sessions[:len(sessions)-1]
	for i := 0; i < len(sessions); i++ {
		sessions[i] = strings.Trim(sessions[i], "\n ")
		// println(sessions[i])
	}
}