package api

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// getHistory godoc
// @Summary Get history
// @Description Obtain user reading history
// @Tags history
// @Produce json
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 401
// @Failure 500
// @Router /history [get]
func (s *Server) getHistory(c *gin.Context) {
	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "init-data not found",
		})
		return
	}

	userId := uint64(initData.User.ID)
	history, err := s.store.GetHistory(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var res []types.GetHistoryResponse
	for i := range history {
		var buf bytes.Buffer
		title := strings.Split(strings.Split(history[i].Manga, ".")[0], "-")
		for _, t := range title {
			buf.WriteString(strings.Title(t) + " ")
		}

		res = append(res, types.GetHistoryResponse{
			Id:     history[i].Id,
			Manga:  buf.String(),
			ReadAt: history[i].ReadAt.Format(time.DateTime),
		})
	}

	c.JSON(http.StatusOK, res)
}

// removeHistory godoc
// @Summary Remove from history
// @Description Delete record from the history
// @Tags history
// @Produce plain
// @Param id query string true "Requested id"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 401
// @Failure 500
// @Router /history [get]
func (s *Server) removeHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "init-data not found",
		})
		return
	}

	userId := uint64(initData.User.ID)
	if err := s.store.RemoveHistory(userId, id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
