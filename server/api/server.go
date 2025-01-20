package api

import (
	_ "github.com/dxtym/yomu/server/api/docs"
	"github.com/dxtym/yomu/server/api/middleware"
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Server struct {
	store  *db.Store
	rdb    *redis.Client
	router *gin.Engine
	config *internal.Config
	scrape *internal.Scrape
}

// @title Yomu API
// @version 1.0
// @description Yomu is a free manga reader Telegram mini app.
// @license.name MIT
// @license.url https://mit-license.org/
// @host localhost:8080
// @basePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @descriprion Enter from Telegram to send your TMA hashkey.
func NewServer(store *db.Store, rdb *redis.Client, scrape *internal.Scrape, config *internal.Config) *Server {
	server := &Server{
		rdb:    rdb,
		store:  store,
		config: config,
		scrape: scrape,
	}
	server.setUpRouting()
	return server
}

func (s *Server) setUpRouting() {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	r := router.Group("/api/v1")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := r.Use(middleware.AuthMiddleware(s.config.BotToken))

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
