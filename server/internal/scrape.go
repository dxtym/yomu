package internal

import (
	"regexp"
	"strings"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gocolly/colly"
)

type Scrape struct {
	url   string
	colly *colly.Collector
}

func NewScrape(url string) *Scrape {
	return &Scrape{
		url: url,
		colly: colly.NewCollector(
			colly.AllowURLRevisit(),
		),
	}
}

func (s *Scrape) GetManga(manga string, res *types.GetMangaResponse) {
	s.colly.OnHTML("#single_book > div.text > div > h1", func(h *colly.HTMLElement) { res.Title = h.Text })

	s.colly.OnHTML("#single_book > div.media > div > img", func(h *colly.HTMLElement) { res.CoverImage = h.Attr("src") })

	s.colly.OnHTML("#single_book > div.summary > p", func(h *colly.HTMLElement) { res.Description = h.Text })

	s.colly.OnHTML("#single_book > div.chapters > table > tbody > tr > td > div > a", func(h *colly.HTMLElement) {
		chapter := struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}{
			Name: h.Text,
			Url:  h.Attr("href"),
		}
		res.Chapters = append(res.Chapters, chapter)
	})

	s.colly.Visit(s.url + "manga/" + manga)
}

func (s *Scrape) SearchManga(title string, res *[]types.SearchMangaResponse) {
	var mangas []string
	s.colly.OnHTML("#book_list > div > div.text > h3 > a", func(e *colly.HTMLElement) {
		url := strings.Split(e.Attr("href"), "/")[4]
		mangas = append(mangas, url)
	})

	var images []string
	s.colly.OnHTML("#book_list > div > div.media > div.wrap_img > a > img", func(e *colly.HTMLElement) { images = append(images, e.Attr("src")) })

	s.colly.Visit(s.url + "?search=" + title)

	for i := range mangas {
		*res = append(*res, types.SearchMangaResponse{
			Manga:      mangas[i],
			CoverImage: images[i],
		})
	}
}

func (s *Scrape) GetChapter(manga string, chapter string, res *types.GetChapterResponse) {
	s.colly.OnResponse(func(r *colly.Response) {
		body := string(r.Body)
		re := regexp.MustCompile(`var\sthzq=\[(.*?)\];`)
		urls := regexp.MustCompile(`https:\/\/([^\']+)`)

		match := re.FindString(body)
		res.PageUrls = urls.FindAllString(match, -1)
	})

	s.colly.Visit(s.url + "manga/" + manga + "/" + chapter)
}
