package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

const _initDataKey = "init-data"

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, ngrok-skip-browser-warning")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func withInitData(c context.Context, initData initdata.InitData) context.Context {
	return context.WithValue(c, _initDataKey, initData)
}

func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.GetHeader("Authorization"), " ")
		if len(auth) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		atype := auth[0]
		adata := auth[1]

		switch atype {
		case "tma":
			if err := initdata.Validate(adata, token, time.Hour); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			data, err := initdata.Parse(adata)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.Set("init-data", data)
		}

		c.Next()
	}
}
