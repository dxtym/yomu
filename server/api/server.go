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

func NewServer(store *db.Store, rdb *redis.Client, scrape *internal.Scrape, config *internal.Config) *Server {
	server := &Server{store: store}

	server.rdb = rdb
	server.config = config
	server.scrape = scrape
	server.setUpRouting()

	return server
}

func (s *Server) setUpRouting() {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	r := router.Group("/api/v1")

	auth := r.Use(middleware.AuthMiddleware(s.config.BotToken))
	auth.POST("/user", s.createUser)
	auth.GET("/search", s.searchManga)
	auth.GET("/manga/:manga", s.getManga)
	auth.GET("/history", s.getHistory)
	auth.DELETE("/history", s.removeHistory)
	auth.POST("/library", s.addLibrary)
	auth.GET("/library", s.getLibrary)
	auth.DELETE("/library", s.removeLibrary)
	auth.GET("/progress", s.getProgress)
	auth.PUT("/progress", s.updateProgress)
	auth.GET("/chapter/:manga/:chapter", s.getChapter)

	s.router = router
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
