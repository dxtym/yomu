package handlers

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db/models"
	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create user
// @Description Register user via Telegram data
// @Tags user
// @Accept json
// @Produce json
// @Param data body types.CreateUserRequest true "Requested user"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success 201
// @Failure 400
// @Failure 500
// @Router /user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req types.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrResponse(err))
		return
	}

	user := &models.User{
		UserId:    req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
	}
	if err := h.db.CreateUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse(err))
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
