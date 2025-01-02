package api

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type GetChapterResponse struct {
	PageUrls []string `json:"page_urls"`
}

func (s *Server) getChapter(c *gin.Context) {
	url := c.Param("url")
	id := c.Param("id")

	var res GetChapterResponse
	s.colly.OnResponse(func(r *colly.Response) {
		body := string(r.Body)
		re := regexp.MustCompile(`var\sthzq=\[(.*?)\];`)
		urls := regexp.MustCompile(`https:\/\/([^\']+)`)
		match := re.FindString(body)
		res.PageUrls = urls.FindAllString(match, -1)
		log.Printf("found: %s\n", res.PageUrls)
	})

	s.colly.Visit(s.config.ApiUrl + "manga/" + url + "/" + id)

	c.JSON(http.StatusOK, res)
}
