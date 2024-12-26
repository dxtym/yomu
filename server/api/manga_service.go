package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

type SearchMangaQuery struct {
	Page struct {
		Media []struct {
			Id         uint
			CoverImage struct {
				Large string
			}
		}
	}
}

type SearchMangaResponse struct {
	MangaId    uint   `json:"manga_id"`
	CoverImage string `json:"cover_image"`
}

func (s *Server) searchManga(c *gin.Context) {
	title := c.Query("title")

	var mangaQuery SearchMangaQuery
	query := "{ Page (page: 1, perPage: 10) { media (type: MANGA, isAdult: false, search: %q) { id coverImage { large } } } }"
	req := graphql.NewRequest(fmt.Sprintf(query, title))
	if err := s.client.Run(c, req, &mangaQuery); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var resp []SearchMangaResponse
	for _, media := range mangaQuery.Page.Media {
		resp = append(resp, SearchMangaResponse{
			MangaId:    media.Id,
			CoverImage: media.CoverImage.Large,
		})
	}

	c.JSON(http.StatusOK, resp)
}
