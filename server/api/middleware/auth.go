package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.GetHeader("Authorization"), " ")
		if len(auth) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized access",
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
