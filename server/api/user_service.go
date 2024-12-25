package api

import (
	"net/http"
	"time"

	"github.com/dxtym/yomu/server/db"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	UserId    uint   `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
}

type CreateUserResponse struct {
	UserId    uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Token     string `json:"token"`
}

func (s *Server) createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := &db.User{
		UserId:    req.UserId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
	}
	if err := s.store.CreateUser(user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := s.token.CreateToken(user.UserId, 24*time.Hour)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res := CreateUserResponse{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Token:     token,
	}

	c.JSON(http.StatusCreated, res)
}
