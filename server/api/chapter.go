package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
)

func (s *Server) getChapter(c *gin.Context) {
	manga := c.Param("manga")
	chapter := c.Param("chapter")

	var res types.GetChapterResponse
	s.scrape.GetChapter(s.config.ApiUrl, manga, chapter, &res)

	c.JSON(http.StatusOK, res)
}
