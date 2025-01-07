package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (s *Server) getManga(c *gin.Context) {
	manga := c.Param("url")

	var res types.GetMangaResponse
	val, err := s.rdb.Get(c, manga).Result()
	if err == redis.Nil {
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

		c.JSON(http.StatusOK, res)
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	} else {
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, res)
		return
	}
}

func (s *Server) searchManga(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusNoContent, gin.H{"message": "empty title"})
		return
	}

	var res []types.SearchMangaResponse
	s.scrape.SearchManga(s.config.ApiUrl, title, &res)

	c.JSON(http.StatusOK, res)
}
