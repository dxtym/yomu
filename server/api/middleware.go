package api

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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func withInitData(c context.Context, initData initdata.InitData) context.Context {
	return context.WithValue(c, _initDataKey, initData)
}

func authMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authParts := strings.Split(c.GetHeader("authorization"), " ")
		if len(authParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		authType := authParts[0]
		authData := authParts[1]

		switch authType {
		case "tma":
			if err := initdata.Validate(authData, token, time.Hour); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			initData, err := initdata.Parse(authData)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.Request = c.Request.WithContext(
				withInitData(c.Request.Context(), initData),
			)
		}

		c.Next()
	}
}
