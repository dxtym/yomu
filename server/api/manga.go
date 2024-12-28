package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type SearchMangaResponse struct {
	MangaUrl   string `json:"manga_url"`
	CoverImage string `json:"cover_image"`
}

func (s *Server) searchManga(c *gin.Context) {
	title := c.Query("title")

	var urls []string
	s.colly.OnHTML("#book_list > div > div.text > h3 > a", func(e *colly.HTMLElement) {
		urls = append(urls, e.Attr("href"))
		log.Printf("found: %q -> %s\n", e.Text, e.Attr("href"))
	})

	var images []string
	s.colly.OnHTML("#book_list > div > div.media > div.wrap_img > a > img", func(e *colly.HTMLElement) {
		images = append(images, e.Attr("src"))
		log.Printf("found: %s\n", e.Attr("src"))
	})

	s.colly.OnRequest(func(r *colly.Request) {
		log.Printf("visiting: %s", r.URL.String())
	})
	s.colly.Visit(s.config.ApiUrl + "?search=" + title)

	var res []SearchMangaResponse
	for i := range urls {
		res = append(res, SearchMangaResponse{
			MangaUrl:   urls[i],
			CoverImage: images[i],
		})
	}

	c.JSON(http.StatusOK, res)
}
