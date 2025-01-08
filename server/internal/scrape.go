package internal

import (
	"log"
	"regexp"
	"strings"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gocolly/colly"
)

type Scrape struct {
	colly *colly.Collector
}

func NewScrape() *Scrape {
	return &Scrape{
		colly: colly.NewCollector(
			colly.AllowURLRevisit(),
		),
	}
}

func (s *Scrape) GetManga(url string, manga string, res *types.GetMangaResponse) {
	s.colly.OnHTML("#single_book > div.text > div > h1", func(h *colly.HTMLElement) {
		res.Title = h.Text
		log.Printf("found: %s\n", h.Text)
	})

	s.colly.OnHTML("#single_book > div.media > div > img", func(h *colly.HTMLElement) {
		res.CoverImage = h.Attr("src")
		log.Printf("found: %s\n", h.Attr("src"))
	})

	s.colly.OnHTML("#single_book > div.summary > p", func(h *colly.HTMLElement) {
		res.Description = h.Text
		log.Printf("found: %s\n", h.Text)
	})

	s.colly.OnHTML("#single_book > div.chapters > table > tbody > tr > td > div > a", func(h *colly.HTMLElement) {
		chapter := struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}{
			Name: h.Text,
			Url:  h.Attr("href"),
		}
		res.Chapters = append(res.Chapters, chapter)
		log.Printf("found: %s -> %s\n", h.Text, h.Attr("href"))
	})

	s.colly.Visit(url + "manga/" + manga)
}

func (s *Scrape) SearchManga(url string, title string, res *[]types.SearchMangaResponse) {
	var mangas []string
	s.colly.OnHTML("#book_list > div > div.text > h3 > a", func(e *colly.HTMLElement) {
		url := strings.Split(e.Attr("href"), "/")[4]
		mangas = append(mangas, url)
		log.Printf("found: %q -> %s\n", e.Text, e.Attr("href"))
	})

	var images []string
	s.colly.OnHTML("#book_list > div > div.media > div.wrap_img > a > img", func(e *colly.HTMLElement) {
		images = append(images, e.Attr("src"))
		log.Printf("found: %s\n", e.Attr("src"))
	})

	s.colly.Visit(url + "?search=" + title)

	for i := range mangas {
		*res = append(*res, types.SearchMangaResponse{
			Manga:      mangas[i],
			CoverImage: images[i],
		})
	}
}

func (s *Scrape) GetChapter(url string, manga string, chapter string, res *types.GetChapterResponse) {
	s.colly.OnResponse(func(r *colly.Response) {
		body := string(r.Body)
		re := regexp.MustCompile(`var\sthzq=\[(.*?)\];`)
		urls := regexp.MustCompile(`https:\/\/([^\']+)`)
		
		match := re.FindString(body)
		res.PageUrls = urls.FindAllString(match, -1)
		log.Printf("found: %s\n", res.PageUrls)
	})

	s.colly.Visit(url + "manga/" + manga + "/" + chapter)
}
