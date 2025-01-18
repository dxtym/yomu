package api

import (
	"net/http"
	"time"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// updateProgress godoc
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
func (s *Server) updateProgress(c *gin.Context) {
	var req types.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
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
	progress := &db.Progress{
		UserId:   userId,
		Manga:    req.Manga,
		Chapter:  req.Chapter,
		Page:     req.Page,
		UpdateAt: time.Now(),
	}
	if err := s.store.UpdateProgress(progress); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	history := &db.History{
		UserId: userId,
		Manga:  req.Manga,
		ReadAt: time.Now(),
	}
	if err := s.store.AddHistory(history); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// getProgress godoc
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
func (s *Server) getProgress(c *gin.Context) {
	manga := c.Query("manga")
	chapter := c.Query("chapter")

	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "init-data not found",
		})
		return
	}

	userId := uint64(initData.User.ID)
	page, err := s.store.GetProgress(userId, manga, chapter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, page)
}
