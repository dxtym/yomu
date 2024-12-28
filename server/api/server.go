package api

import (
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"

	"github.com/gocolly/colly"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
	token  *internal.Token
	config *internal.Config
	colly  *colly.Collector
}

func NewServer(store *db.Store, token *internal.Token, config *internal.Config) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.Use(CorsMiddleware())

	v1 := router.Group("/api/v1")
	v1.POST("/user", server.createUser)

	// TODO: fix invalid token
	auth := v1.Use(authMiddleware(config.BotToken))
	auth.GET("/library", server.getLibrary)
	auth.GET("/search", server.searchManga)

	server.router = router
	server.token = token
	server.config = config
	server.colly = colly.NewCollector()

	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
