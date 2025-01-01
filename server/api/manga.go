package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type GetMangaResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
}

type SearchMangaResponse struct {
	MangaUrl   string `json:"manga_url"`
	CoverImage string `json:"cover_image"`
}

func (s *Server) getManga(c *gin.Context) {
	url := c.Param("url")

	var res GetMangaResponse
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

	s.colly.Visit(s.config.ApiUrl + "manga/" + url)

	c.JSON(http.StatusOK, res)
}

func (s *Server) searchManga(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusNoContent, gin.H{"message": "empty title"})
		return
	}

	var urls []string
	s.colly.OnHTML("#book_list > div > div.text > h3 > a", func(e *colly.HTMLElement) {
		url := strings.Split(e.Attr("href"), "/")[4]
		urls = append(urls, url)
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
