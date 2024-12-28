package api

import (
	"net/http"

	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Id        int64  `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
}

type CreateUserResponse struct {
	Id    int64  `json:"id"`
	FirstName string `json:"first_name"`
}

func (s *Server) createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := &db.User{
		Id:        req.Id,
		FirstName: req.FirstName,
	}
	if err := s.store.CreateUser(user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, CreateUserResponse{
		Id:        user.Id,
		FirstName: user.FirstName,	
	})
}
