package api

import (
	_ "github.com/dxtym/yomu/server/api/docs"
	"github.com/dxtym/yomu/server/api/handlers"
	"github.com/dxtym/yomu/server/api/middleware"
	"github.com/dxtym/yomu/server/db/store"
	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	db     *store.Store
	rdb    *redis.Client
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
func NewServer(db *store.Store, rdb *redis.Client, config *internal.Config) *Server {
	scrape := internal.NewScrape(config.ApiUrl)
	server := &Server{
		db:     db,
		rdb:    rdb,
		scrape: scrape,
		config: config,
	}
	
	// TODO: set up validator
	server.setUpRouting()
	return server
}

func (s *Server) setUpRouting() {
	router := gin.Default()
	handler := handlers.NewHandler(s.db, s.rdb, s.scrape)
	router.Use(middleware.CorsMiddleware())

	r := router.Group("/api/v1")
	r.POST("/user", handler.CreateUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Use(middleware.AuthMiddleware(s.config.BotToken))
	auth.GET("/search", handler.SearchManga)
	auth.GET("/manga/:manga", handler.GetManga)
	auth.GET("/history", handler.GetHistory)
	auth.DELETE("/history", handler.RemoveHistory)
	auth.GET("/library", handler.GetLibrary)
	auth.POST("/library", handler.AddLibrary)
	auth.DELETE("/library", handler.RemoveLibrary)
	auth.GET("/progress", handler.GetProgress)
	auth.PUT("/progress", handler.UpdateProgress)
	auth.GET("/chapter/:manga/:chapter", handler.GetChapter)

	s.router = router
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
