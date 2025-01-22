package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dxtym/yomu/server/api/handlers"
	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.GetHeader("Authorization"), " ")
		if len(auth) != 2 {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				handlers.ErrResponse(errors.New("incorrect auth format")),
			)
			return
		}

		kind := auth[0]
		data := auth[1]

		switch kind {
		case "tma":
			if err := initdata.Validate(data, token, time.Hour); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.ErrResponse(err))
				return
			}

			initData, err := initdata.Parse(data)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, handlers.ErrResponse(err))
				return
			}

			c.Set("init-data", initData)
		}

		c.Next()
	}
}
