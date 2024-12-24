package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	UserId uint `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
}

func (s *Server) createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := &db.User{
		UserId: req.UserId,
		FirstName: req.FirstName,
		LastName: req.LastName,
		UserName: req.UserName,
	}
	if err := s.store.CreateUser(user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	
	c.JSON(http.StatusCreated, user)
}