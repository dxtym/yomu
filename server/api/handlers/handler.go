package handlers

import (
	"github.com/dxtym/yomu/server/db/store"
	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	db     *store.Store
	rdb    *redis.Client
	scrape *internal.Scrape
}

func NewHandler(db *store.Store, rdb *redis.Client, scrape *internal.Scrape) *Handler {
	return &Handler{
		db:     db,
		rdb:    rdb,
		scrape: scrape,
	}
}

func ErrResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
