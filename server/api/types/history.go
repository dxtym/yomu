package types

type GetHistoryResponse struct {
	Id        int64  `json:"id"`
	Manga     string `json:"manga"`
	UpdatedAt string `json:"updated_at"`
}
