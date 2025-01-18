package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
)

// createUser godoc
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
func (s *Server) createUser(c *gin.Context) {
	var req types.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user := &db.User{
		Id:        req.Id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
	}
	if err := s.store.CreateUser(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
