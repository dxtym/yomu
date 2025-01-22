package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type HistoryHandler interface {
	GetHistory(c *gin.Context)
	RemoveHistory(c *gin.Context)
}

// GetHistory godoc
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
func (h *Handler) GetHistory(c *gin.Context) {
	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			ErrResponse(errors.New("init-data not found")),
		)
		return
	}

	history, err := h.db.GetHistory(initData.User.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
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
			Id:        history[i].UserId,
			Manga:     buf.String(),
			UpdatedAt: history[i].UpdatedAt.Format(time.DateOnly),
		})
	}

	c.JSON(http.StatusOK, res)
}

// RemoveHistory godoc
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
func (h *Handler) RemoveHistory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			ErrResponse(errors.New("init-data not fond")),
		)
		return
	}

	if err := h.db.RemoveHistory(id, initData.User.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.Status(http.StatusOK)
}
