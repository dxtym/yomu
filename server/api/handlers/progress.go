package handlers

import (
	"errors"
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db/models"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// UpdateProgress godoc
// @Summary Update progress
// @Description Renew current manga reading status
// @Tags progress
// @Accept json
// @Produce plain
// @Param data body types.UpdateProgressRequest true "Requested progress"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /progress [put]
func (h *Handler) UpdateProgress(c *gin.Context) {
	var req types.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrResponse(err))
		return
	}

	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			ErrResponse(errors.New("init-data not found")),
		)
		return
	}

	progress := &models.Progress{
		UserId:  initData.User.ID,
		Manga:   req.Manga,
		Chapter: req.Chapter,
		Page:    req.Page,
	}
	if err := h.db.UpdateProgress(progress); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	history := &models.History{
		UserId: initData.User.ID,
		Manga:  req.Manga,
	}
	if err := h.db.AddHistory(history); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.Status(http.StatusOK)
}

// GetProgress godoc
// @Summary Get progress
// @Description Obtain user progress on chapter
// @Tags progress
// @Produce json
// @Param manga query string true "Requested progress"
// @Param chapter query string true "Requested chapter"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 401
// @Failure 500
// @Router /progress [get]
func (h *Handler) GetProgress(c *gin.Context) {
	manga := c.Query("manga")
	chapter := c.Query("chapter")

	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			ErrResponse(errors.New("init-data not found")),
		)
		return
	}

	page, err := h.db.GetProgress(initData.User.ID, manga, chapter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, page)
}
