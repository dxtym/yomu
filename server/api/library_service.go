package api

import (
	"fmt"
	"net/http"

	"github.com/dxtym/yomu/server/internal"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

type GetLibraryQuery struct {
	Media struct {
		Id         int
		CoverImage struct {
			Large string
		}
	}
}

type GetLibraryResponse struct {
	Library    uint   `json:"library"`
	CoverImage string `json:"cover_image"`
}

func (s *Server) getLibrary(c *gin.Context) {
	// TODO: change from user_id to telegram_id
	user_id := c.MustGet("user_id").(*internal.Claim).UserId

	library, err := s.store.GetLibrary(user_id)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// TODO: change to concurrent pattern
	var libraryQuery []GetLibraryQuery
	for _, id := range library {
		query := fmt.Sprintf("{ Media (id : %d, type: MANGA) { id coverImage { large } } }", id)
		req := graphql.NewRequest(query)

		var res GetLibraryQuery
		if err := s.client.Run(c, req, &res); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		libraryQuery = append(libraryQuery, res)
	}

	var resp []GetLibraryResponse
	for i := range libraryQuery {
		resp = append(resp, GetLibraryResponse{
			Library:    library[i],
			CoverImage: libraryQuery[i].Media.CoverImage.Large,
		})
	}

	c.JSON(http.StatusOK, resp)
}
