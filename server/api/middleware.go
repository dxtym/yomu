package api

import (
	"net/http"
	"strings"

	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authMiddleware(token *internal.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if fields[0] != "Bearer" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		authToken := fields[1]
		userId, err := token.VerifyToken(authToken)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
