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
	config *internal.Config
	colly  *colly.Collector
}

func NewServer(store *db.Store, config *internal.Config) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.Use(internal.CorsMiddleware())

	v1 := router.Group("/api/v1")
	v1.POST("/user", server.createUser)

	auth := v1.Use(internal.AuthMiddleware(config.BotToken))
	auth.GET("/library", server.getLibrary)
	auth.GET("/search", server.searchManga)
	auth.GET("/manga/:url", server.getManga)
	auth.GET("/chapter/:url/:id", server.getChapter)

	server.router = router
	server.config = config
	server.colly = colly.NewCollector()

	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
