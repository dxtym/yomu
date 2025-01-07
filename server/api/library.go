package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *Server) addLibrary(c *gin.Context) {
	var req types.AddLibraryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	initData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "init-data not found",
		})
		return
	}

	record := &db.Library{
		UserId:     uint(initData.User.ID),
		Manga:      req.Manga,
		CoverImage: req.CoverImage,
	}
	if err := s.store.AddLibrary(record); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}

func (s *Server) getLibrary(c *gin.Context) {
	intiData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "init-data not found",
		})
	}

	library, err := s.store.GetLibrary(uint(intiData.User.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var res []types.GetLibraryResponse
	for _, record := range library {
		res = append(res, types.GetLibraryResponse{
			MangaUrl:   record.Manga,
			CoverImage: record.CoverImage,
		})
	}

	c.JSON(http.StatusOK, res)
}
