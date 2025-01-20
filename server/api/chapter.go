package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
)

// getChapter godoc
// @Summary Get schapter
// @Description Scrape page urls of the chapter
// @Tags chapter
// @Produce json
// @Param manga path string true "Requested manga"
// @Param chapter path string true "Requested chapter"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 500
// @Router /chapter/{manga}/{chapter} [get]
func (s *Server) getChapter(c *gin.Context) {
	manga := c.Param("manga")
	chapter := c.Param("chapter")

	var res types.GetChapterResponse
	s.scrape.GetChapter(s.config.ApiUrl, manga, chapter, &res)

	if len(res.PageUrls) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "no pages found",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
