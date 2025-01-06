package server

import (
	"github.com/dxtym/yomu/server/api/middleware"
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	store  *db.Store
	rdb    *redis.Client
	router *gin.Engine
	config *internal.Config
	scrape *internal.Scrape
}

func NewServer(
	store *db.Store,
	rdb *redis.Client,
	scrape *internal.Scrape,
	config *internal.Config,
) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	v1 := router.Group("/api/v1")
	v1.POST("/user", server.createUser)

	auth := v1.Use(middleware.AuthMiddleware(config.BotToken))
	auth.GET("/search", server.searchManga)
	auth.GET("/manga/:url", server.getManga)
	auth.GET("/chapter/:url/:id", server.getChapter)
	auth.POST("/library", server.addLibrary)
	auth.GET("/library", server.getLibrary)

	server.rdb = rdb
	server.router = router
	server.config = config
	server.scrape = scrape

	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
