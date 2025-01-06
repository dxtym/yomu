package server

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) getChapter(c *gin.Context) {
	url := c.Param("url")
	id := c.Param("id")

	var res types.GetChapterResponse
	s.scrape.GetChapter(s.config.ApiUrl, url, id, &res)

	c.JSON(http.StatusOK, res)
}
