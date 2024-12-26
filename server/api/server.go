package api

import (
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
	token  *internal.Token
	client *graphql.Client
}

func NewServer(store *db.Store, token *internal.Token, client *graphql.Client) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.Use(CorsMiddleware())

	v1 := router.Group("/api/v1")
	v1.POST("/user", server.createUser)

	// TODO: fix invalid token
	auth := v1.Group("/").Use(authMiddleware(token))
	auth.GET("/library", server.getLibrary)
	auth.GET("/search", server.searchManga)

	server.router = router
	server.token = token
	server.client = client

	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
