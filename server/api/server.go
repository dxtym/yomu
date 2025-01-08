package api

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

	r := router.Group("/api/v1")
	auth := r.Use(middleware.AuthMiddleware(config.BotToken))

	auth.POST("/user", server.createUser)
	auth.GET("/search", server.searchManga)
	auth.GET("/manga/:manga", server.getManga)
	auth.GET("/chapter/:manga/:chapter", server.getChapter)
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
