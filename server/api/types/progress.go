package types

type UpdateProgressRequest struct {
	Manga   string `json:"manga" binding:"required"`
	Chapter string `json:"chapter" binding:"required"`
	Page    int64 `json:"page" binding:"required"`
}
