package api

import (
	"log"
	"net/http"

	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type GetLibraryResponse struct {
	MangaUrl   string `json:"manga_url"`
	CoverImage string `json:"cover_image"`
}

func (s *Server) getLibrary(c *gin.Context) {
	_ = c.MustGet("user_id").(*internal.Claim).UserId

	var res []GetLibraryResponse
	s.colly.OnHTML("#single_book > div.d-cell-medium.media > div > img", func(e *colly.HTMLElement) {
		mangaUrl := e.Request.URL.String()
		res = append(res, GetLibraryResponse{
			MangaUrl:   mangaUrl,
			CoverImage: e.Attr("src"),
		})
		log.Printf("found: %s\n", e.Attr("src"))
	})

	s.colly.OnRequest(func(r *colly.Request) {
		log.Printf("visiting: %s", r.URL.String())
	})

	c.JSON(http.StatusOK, res)
}
