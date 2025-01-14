package api

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *Server) getHistory(c *gin.Context) {
	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "init-data not found",
		})
		return
	}

	userId := uint(initData.User.ID)
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
			Manga:  buf.String(),
			ReadAt: history[i].ReadAt.Format(time.DateTime),
		})
	}

	c.JSON(http.StatusOK, res)
}
