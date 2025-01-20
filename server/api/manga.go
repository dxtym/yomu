package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// getManga godoc
// @Summary Get manga
// @Description Obtain details about manga
// @Tags manga
// @Produce json
// @Param manga path string true "Requested manga"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 500
// @Router /manga/{manga} [get]
func (s *Server) getManga(c *gin.Context) {
	manga := c.Param("manga")

	var res types.GetMangaResponse
	val, err := s.rdb.Get(c, manga).Result()

	switch err {
	case redis.Nil:
		s.scrape.GetManga(s.config.ApiUrl, manga, &res)
		data, err := json.Marshal(res)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		err = s.rdb.Set(c, manga, data, time.Hour).Err()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	case nil:
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// searchManga godoc
// @Summary Search manga
// @Description Search for manga by title
// @Tags manga
// @Produce json
// @Param query query string true "Requested title"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 204
// @Failure 500
// @Router /search [get]
func (s *Server) searchManga(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusNoContent, gin.H{"message": "empty query"})
		return
	}

	var res []types.SearchMangaResponse
	s.scrape.SearchManga(s.config.ApiUrl, query, &res)

	if len(res) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "no manga found",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
