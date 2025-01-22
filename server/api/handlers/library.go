package handlers

import (
	"errors"
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db/models"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type LibraryHandler interface {
	AddLibrary(c *gin.Context)
	GetLibrary(c *gin.Context)
	RemoveLibrary(c *gin.Context)
}

// AddLibrary godoc
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
func (h *Handler) AddLibrary(c *gin.Context) {
	var req types.AddLibraryRequest
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

	record := &models.Library{
		UserId:     initData.User.ID,
		Manga:      req.Manga,
		CoverImage: req.CoverImage,
	}
	if err := h.db.AddLibrary(record); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.Status(http.StatusCreated)
}

// GetLibrary godoc
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
func (h *Handler) GetLibrary(c *gin.Context) {
	intiData, ok := c.MustGet("init-data").(initdata.InitData)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			ErrResponse(errors.New("init-data not found")),
		)
	}

	library, err := h.db.GetLibrary(intiData.User.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
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

// RemoveLibrary godoc
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
func (h *Handler) RemoveLibrary(c *gin.Context) {
	var req types.RemoveLibraryRequest
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

	err := h.db.RemoveLibrary(initData.User.ID, req.Manga)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.Status(http.StatusOK)
}
