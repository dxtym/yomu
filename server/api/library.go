package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

// addLibrary godoc
// @Summary Add to library
// @Description Create new manga in the library
// @Tags library
// @Accept json
// @Produce plain
// @Param data body types.AddLibraryRequest true "Requested add"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /library [post]
func (s *Server) addLibrary(c *gin.Context) {
	var req types.AddLibraryRequest
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

	record := &db.Library{
		UserId:     uint64(initData.User.ID),
		Manga:      req.Manga,
		CoverImage: req.CoverImage,
	}
	if err := s.store.AddLibrary(record); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

// getLibrary godoc
// @Summary Get from library
// @Description Obtain mangas in the library
// @Tags library
// @Produce json
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 401
// @Failure 500
// @Router /library [get]
func (s *Server) getLibrary(c *gin.Context) {
	intiData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "init-data not found",
		})
	}

	library, err := s.store.GetLibrary(uint64(intiData.User.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var res []types.GetLibraryResponse
	for _, record := range library {
		res = append(res, types.GetLibraryResponse{
			Manga:      record.Manga,
			CoverImage: record.CoverImage,
		})
	}

	c.JSON(http.StatusOK, res)
}

// removeLibrary godoc
// @Summary Remove from library
// @Description Delete manga from the library
// @Tags library
// @Accept json
// @Produce plain
// @Param data body types.RemoveLibraryRequest true "Requested delete"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /library [delete]
func (s *Server) removeLibrary(c *gin.Context) {
	var req types.RemoveLibraryRequest
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
	err := s.store.RemoveLibrary(userId, req.Manga)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
