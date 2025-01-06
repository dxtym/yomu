package server

import (
	"net/http"

	"github.com/dxtym/yomu/server/api/types"
	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
)

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
