package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type MangaHandler interface {
	GetManga(c *gin.Context)
	SearchManga(c *gin.Context)
}

// GetManga godoc
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
func (h *Handler) GetManga(c *gin.Context) {
	manga := c.Param("manga")

	var res types.GetMangaResponse
	val, err := h.rdb.Get(c, manga).Result()

	switch err {
	case redis.Nil:
		h.scrape.GetManga(manga, &res)
		data, err := json.Marshal(res)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
			return
		}

		err = h.rdb.Set(c, manga, data, time.Hour).Err()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
			return
		}
	case nil:
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
			return
		}
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

// SearchManga godoc
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
func (h *Handler) SearchManga(c *gin.Context) {
	query := c.Query("query")

	var res []types.SearchMangaResponse
	h.scrape.SearchManga(query, &res)
	if len(res) == 0 {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, 
			ErrResponse(errors.New("no manga found")),
		)
		return
	}

	c.JSON(http.StatusOK, res)
}
